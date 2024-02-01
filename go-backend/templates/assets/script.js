document.addEventListener("DOMContentLoaded", function() {
    anime({
        targets: 'body',
        opacity: 1,
        easing: 'easeInOutQuad',
        duration: 1000,
        delay: 500
    });
});

document.addEventListener('DOMContentLoaded', function () {
    var contactForm = document.getElementById('contact-form');

    contactForm.addEventListener('submit', function (event) {
        event.preventDefault();
        var formData = new FormData(contactForm);

        fetch('http://localhost:8080/callback', {
            method: 'POST',
            body: formData
        })
            .then(response => response.json())
            .then(data => {
                console.log(data);

                if (data.response.message === "success 2xx") {
                    alert("Ваша заявка успешно отправлена!");
                } else {
                    alert("Произошла ошибка при отправке заявки. Пожалуйста, попробуйте еще раз.");
                }
            })
            .catch(error => {
                console.error('Error:', error);
            });
    });
});

