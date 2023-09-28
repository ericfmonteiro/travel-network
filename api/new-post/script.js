function handleSubmit(event) {
    event.preventDefault();

    var form = document.getElementById("registrationForm");
    var formData = {
        title: form.querySelector("#title").value,
        content: form.querySelector("#content").value,
    };


    var queryParams = new URLSearchParams(window.location.search);
    var username = queryParams.get("username");

    getUserByName(username)
        .then(user => {
            return createPost(formData, user)
        })
        .then(random => {
            alert("Post created successfully!");
            form.reset();
            window.location.href = '../feed/feed.html?username=' + username;
        })
        .catch(error => {
            console.log('Error:', error)
        })
}

document.getElementById("registrationForm").addEventListener("submit", handleSubmit);

function getUserByName(username) {
    return fetch("http://localhost:8080/users/name/" + username, {
        method: "GET",
        headers: {
            "Content-Type": "application/json"
        }
    })
        .then(response => {
            if (response.ok) {
                console.log(response)
                return response.json()
            } else {
                alert("Error getting user. Please try again.");
            }
        })
        .catch(error => {
            console.error("Error:", error);
            alert("An error occurred. Please try again later.");
        });
}

function createPost(formData, user) {
    formData.userId = user.id

    console.log(formData)
    console.log(user)

    var jsonData = JSON.stringify(formData);


    fetch("http://localhost:8080/posts/create", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: jsonData
    })
        .then(response => {
            if (response.ok) {
                console.log(response)
                return response.json()
            } else {
                alert("Error creating user. Please try again.");
            }
        })
        .catch(error => {
            console.error("Error:", error);
            alert("An error occurred. Please try again later.");
        });
}