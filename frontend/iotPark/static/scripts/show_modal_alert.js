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
    try {
        var response = JSON.parse(event.detail.xhr.response);

        if (response.hasOwnProperty('error')) {
            showAlert(response.error, 'danger');
        } else {
            showAlert(response, 'success');
        }
    } catch (e) {
        // In case of parsing error, show a default error message
        showAlert('An error occurred', 'danger');
    }
});
