function validatePassword() {
    var password = document.getElementById("senha").value;
    var confirmPassword = document.getElementById("confirm_senha").value;

    if (password !== confirmPassword) {
        alert("As senhas não coincidem. Por favor, digite a mesma senha nos campos de senha e confirmar senha.");
        return false;
    }
    return true;
}

function handleSubmit(event) {
    event.preventDefault(); 

    if (!validatePassword()) {
        return; 
    }

    var form = document.getElementById("registrationForm");
    var formData = {
        name: form.querySelector("#nome").value,
        email: form.querySelector("#email").value,
        numtrips: form.querySelector("#num_viagens").value,
        bio: form.querySelector("#biografia").value,
        cpf: form.querySelector("#cpf").value,
        password: form.querySelector("#senha").value,
    };

    var jsonData = JSON.stringify(formData);

    const username = form.querySelector("#nome").value;

    fetch(form.action, {
        method: "POST",
        headers: {
            "Content-Type": "application/json" 
        },
        body: jsonData
    })
        .then(response => {
            if (response.ok) {
                alert("User created successfully!");
                form.reset();
                window.location.href = '../feed/feed.html?username=' + username;
            } else {
                alert("Error creating user. Please try again.");
            }
        })
        .catch(error => {
            console.error("Error:", error);
            alert("An error occurred. Please try again later.");
        });
}

document.getElementById("registrationForm").addEventListener("submit", handleSubmit);