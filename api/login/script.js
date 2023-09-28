const loginForm = document.getElementById('login-form');
loginForm.addEventListener('submit', function (event) {
    event.preventDefault();

    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    // Simulate authentication (replace with actual authentication logic)
    // if (username === 'user' && password === 'password') {
    //     alert('Login successful!');
    //     window.location.href = '../feed/feed.html?username=' + username + '&password=' + password;
    //     // Redirect to another page or perform desired actions
    // } else {
    //     alert('Invalid credentials. Please try again.');
    // }

    // const username = "user";
    // const password = "password";
    const credentials = btoa(username + ":" + password); // Encode to base64

    fetch('http://localhost:8080', {
        method: 'GET',
        headers: {
            "Origin": "http://locahost:5500", // Ensure this matches the allowed origin on the server
            "Authorization": "Basic " + credentials,
        }
    })
        .then(response => {
            if (response.status == 202) {
                console.log(response)
                window.location.href = '../feed/feed.html?username=' + username;
            } else {
                alert("Invalid credentials")
            }
            return response.json();
        })
        .then((response) => {
            console.log(response);
        })
        .catch(error => {
            console.error('Fetch error:', error);
        });
});

// Create event listener to edit user
var createUserButton = document.getElementById("create-user-button");
createUserButton.addEventListener("click", function () {
    window.location.href = "../new-user/new-user.html";
});