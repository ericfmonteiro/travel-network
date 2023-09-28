function handleSubmit(event) {
    event.preventDefault(); 

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

    var queryParams = new URLSearchParams(window.location.search);
    var oldUserName = queryParams.get("username");

    fetch("http://localhost:8080/users/" + oldUserName + "/edit", {
        method: "PUT",
        headers: {
            "Content-Type": "application/json" 
        },
        body: jsonData
    })
        .then(response => {
            if (response.ok) {
                alert("User edited successfully!");
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