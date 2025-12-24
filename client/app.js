const API_BASE_URL = "http://localhost:8080/api";


let cart = [];
let itemsCatalog = []; 

$(document).ready(function () {
    const path = window.location.pathname;
    const isDashboard = path.includes("dashboard.html");
    if (!isDashboard) {
        if (localStorage.getItem("token")) {
            window.location.href = "dashboard.html";
        }
        $("#login-form").submit(function (e) {
            e.preventDefault();
            const username = $("#username").val();
            const password = $("#password").val();
            $.ajax({
                url: `${API_BASE_URL}/login`,
                method: "POST",
                contentType: "application/json",
                data: JSON.stringify({ username, password }),
                success: function (res) {
                    localStorage.setItem("token", res.token);
                    
                    Swal.fire({
                        icon: 'success',
                        title: 'Login Success',
                        text: 'Redirecting...',
                        timer: 1500,
                        showConfirmButton: false
                    }).then(() => {
                        window.location.href = "dashboard.html";
                    });
                },
                error: function (err) {
                    Swal.fire('Error', err.responseJSON?.error || "Login Failed", 'error');
                }
            });
        });
    } 
    

    else {
        if (!localStorage.getItem("token")) {
            window.location.href = "index.html";
            return;
        }
        loadItems();
        loadSuppliers();
        $("#btn-logout").click(function() {
            localStorage.removeItem("token");
            window.location.href = "index.html";
        });
        $("#btn-add-cart").click(function() {
            const itemId = parseInt($("#item-select").val());
            const qty = parseInt($("#item-qty").val());
            if (!itemId || qty <= 0) {
                Swal.fire('Warning', 'Pilih Item dan masukkan jumlah yang benar', 'warning');
                return;
            }
            const itemData = itemsCatalog.find(i => i.id === itemId);
            const existingItem = cart.find(i => i.item_id === itemId);
            if (existingItem) {
                existingItem.qty += qty; 
            } else {
                cart.push({
                    item_id: itemId,
                    name: itemData.name,
                    price: itemData.price,
                    qty: qty
                });
            }
            renderCart();
            $("#item-qty").val(1); 
        });
        $("#cart-table-body").on("click", ".btn-remove", function() {
            const index = $(this).data("index");
            cart.splice(index, 1); 
            renderCart();
        });
        $("#btn-submit-order").click(function() {
            const supplierId = $("#supplier-select").val();
            if (!supplierId) {
                Swal.fire('Warning', 'Pilih Supplier terlebih dahulu!', 'warning');
                return;
            }
            const payload = {
                supplier_id: parseInt(supplierId),
                items: cart.map(i => ({
                    item_id: i.item_id,
                    qty: i.qty
                }))
            };
            apiRequest("/v1/purchasing", "POST", payload, 
                function(response) {
                    Swal.fire('Success', 'Transaksi Berhasil Disimpan!', 'success');
                    cart = []; 
                    renderCart();
                    loadItems(); 
                    $("#supplier-select").val(""); 
                },
                function(err) {
                    Swal.fire('Error', err.responseJSON?.error || "Transaksi Gagal", 'error');
                }
            );
        });
    }
});


function apiRequest(endpoint, method, data, onSuccess, onError) {
    $.ajax({
        url: API_BASE_URL + endpoint,
        method: method,
        contentType: "application/json",
        headers: {
            "Authorization": "Bearer " + localStorage.getItem("token")
        },
        data: data ? JSON.stringify(data) : null,
        success: onSuccess,
        error: function(xhr) {
            if (xhr.status === 401) {
                alert("Session habis, silakan login ulang.");
                localStorage.removeItem("token");
                window.location.href = "index.html";
            } else {
                if (onError) onError(xhr);
            }
        }
    });
}


function renderCart() {
    const tbody = $("#cart-table-body");
    tbody.empty();
    let total = 0;
    if (cart.length === 0) {
        tbody.html('<tr><td colspan="4" class="text-center text-muted">Cart is empty</td></tr>');
        $("#btn-submit-order").prop("disabled", true);
        $("#grand-total").text("0");
        return;
    }
    cart.forEach((item, index) => {
        const subtotal = item.price * item.qty;
        total += subtotal;  
        tbody.append(`
            <tr>
                <td>${item.name}</td>
                <td>${item.qty}</td>
                <td>Rp ${subtotal.toLocaleString()}</td>
                <td>
                    <button class="btn btn-danger btn-sm btn-remove" data-index="${index}">❌</button>
                </td>
            </tr>
        `);
    });
    $("#grand-total").text("Rp " + total.toLocaleString());
    $("#btn-submit-order").prop("disabled", false);
}


function loadItems() {
    apiRequest("/v1/items", "GET", null, function(data) {
        itemsCatalog = data; 
        const tbody = $("#inventory-table-body");
        tbody.empty();
        data.forEach(item => {
            tbody.append(`
                <tr>
                    <td>${item.id}</td>
                    <td>${item.name}</td>
                    <td>Rp ${item.price.toLocaleString()}</td>
                    <td><span class="badge bg-${item.stock > 0 ? 'success' : 'danger'}">${item.stock}</span></td>
                </tr>
            `);
        });
        const select = $("#item-select");
        select.html('<option value="">-- Choose Item --</option>');
        data.forEach(item => {
            select.append(`<option value="${item.id}">${item.name} (Stock: ${item.stock})</option>`);
        });
    });
}


function loadSuppliers() {
    apiRequest("/v1/suppliers", "GET", null, function(data) {
        const select = $("#supplier-select");
        select.html('<option value="">-- Choose Supplier --</option>');
        data.forEach(s => {
            select.append(`<option value="${s.id}">${s.name}</option>`);
        });
    });
}