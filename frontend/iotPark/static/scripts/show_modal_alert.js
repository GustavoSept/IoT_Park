function showAlert(message, alertType = 'success') {
    var alertHtml = `
        <div class="alert alert-${alertType} alert-dismissible fade show" role="alert">
            ${message}
            <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
        </div>`;

    var alertDiv = document.getElementById('modal-response');
    alertDiv.innerHTML = alertHtml;

    // Automatically close the alert after 3 seconds (3000 milliseconds)
    setTimeout(function() {
        var alert = alertDiv.querySelector('.alert');
        if (alert) {
            new bootstrap.Alert(alert).close();
        }
    }, 3000);
}



document.body.addEventListener('htmx:afterOnLoad', function(event) {
    var response = event.detail.xhr.response;
    showAlert(response, 'info'); // 'info', 'success', 'warning', or 'danger'
});   