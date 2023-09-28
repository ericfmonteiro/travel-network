document.addEventListener("DOMContentLoaded", function () {
    // Get username in query param
    var queryParams = new URLSearchParams(window.location.search);
    var username = queryParams.get("username");
    var userDataElement = document.getElementById("username");
    userDataElement.innerHTML = userDataElement.innerHTML + username;

    // Create event listener to edit user
    var editarUsuarioButton = document.getElementById("edit-user-button");
    editarUsuarioButton.addEventListener("click", function () {
        var queryParams = {
            username: username,
        };

        var url = "../edit-user/edit-user.html?" + Object.entries(queryParams).map(function (pair) {
            return pair.join("=");
        }).join("&");

        window.location.href = url;
    });

    // Create event listener to create post
    var createPostButton = document.getElementById("create-post-button");
    createPostButton.addEventListener("click", function () {
        var queryParams = {
            username: username,
        };

        var url = "../new-post/new-post.html?" + Object.entries(queryParams).map(function (pair) {
            return pair.join("=");
        }).join("&");

        window.location.href = url;
    });

    // Create event listener to logout
    var logoutButton = document.getElementById("logout-button");
    logoutButton.addEventListener("click", function () {
        window.location.href = "../login/index.html";
    });


    // Fetch the feed of the user
    fetch('http://localhost:8080/users/' + username + '/connected-posts', {
        method: "GET"
    })
        .then((response) => response.json())
        .then((data) => {
            const feed = document.querySelector(".feed");

            const fetchLikes = (postId) => {
                return fetch(`http://localhost:8080/posts/likes/${postId}`, {
                    method: "GET"
                }).then((response) => response.json());
            };


            data.forEach((post) => {
                const postElement = document.createElement("div");
                postElement.classList.add("post");

                const titleElement = document.createElement("h2");
                titleElement.classList.add("post-title");
                titleElement.textContent = post.title;

                const textElement = document.createElement("p");
                textElement.classList.add("post-text");
                textElement.textContent = post.content;

                const userElement = document.createElement("p");
                userElement.classList.add("post-user");
                userElement.textContent = `Posted by ${post.userName} on ${post.postDate}`;

                const likesButton = document.createElement("button");
                likesButton.classList.add("post-likes-button");
                likesButton.textContent = "Likes";
                likesButton.dataset.postId = post.id;

                const commentsButton = document.createElement("button");
                commentsButton.classList.add("post-comments-button");
                commentsButton.textContent = "Comments";
                commentsButton.dataset.postId = post.id;

                const commentsContainer = document.createElement("div");
                commentsContainer.classList.add("comments-container");

                const commentInput = document.createElement("input");
                commentInput.classList.add("comment-input");
                commentInput.placeholder = "O que você achou?";

                const commentButton = document.createElement("button");
                commentButton.classList.add("comment-button");
                commentButton.textContent = "Enviar Comentário";

                // Add an event listener to the "Comments" button
                commentButton.addEventListener("click", () => {
                    const commentText = commentInput.value.trim();

                    if (commentText !== "") {
                        const commentData = {
                            postId: post.id,
                            userId: post.userId,
                            content: commentText,
                        };

                        fetch("http://localhost:8080/posts/comments", {
                            method: "POST",
                            headers: {
                                "Content-Type": "application/json"
                            },
                            body: JSON.stringify(commentData)
                        })
                            .then(response => {
                                if (response.ok) {
                                    const commentElement = document.createElement("div");
                                    commentElement.classList.add("comment");
                                    commentElement.textContent = commentText;
                                    commentsContainer.appendChild(commentElement);

                                    commentInput.value = "";
                                } else {
                                    alert("Error adding comment. Please try again.");
                                }
                            })
                            .catch(error => {
                                console.error("Error:", error);
                                alert("An error occurred. Please try again later.");
                            });
                    }
                });


                commentsContainer.appendChild(commentInput);
                commentsContainer.appendChild(commentButton);

                // Fetch likes and comments for the post
                fetchLikes(post.id)
                    .then((likes) => {
                        console.log(likes.length)
                        const likesCount = likes.length;
                        likesButton.textContent = `${likesCount} Likes`;
                    })
                    .catch((error) => {
                        console.error("Error fetching likes:", error);
                    });



                const fetchComments = (postId, containerElement) => {
                    fetch(`http://localhost:8080/posts/comments/${postId}`, {
                        method: "GET"
                    })
                        .then((response) => response.json())
                        .then((comments) => {
                            //containerElement.innerHTML = "";

                            comments.forEach((comment) => {
                                console.log(comment.content)
                                const commentElement = document.createElement("div");
                                commentElement.classList.add("comment");
                                commentElement.textContent = comment.content;
                                console.log(commentElement)
                                containerElement.appendChild(commentElement);
                            });
                        })
                        .catch((error) => {
                            console.error("Error fetching comments:", error);
                        });
                };


                likesButton.addEventListener("click", function () {
                    createLike(post.id, post.userId)
                    var num = likesButton.textContent.replace(/[^0-9]/g, '');
                    num++
                    likesButton.textContent = num + " Likes"
                    likesButton.disabled = true;
                })

                // Add an event listener to the "Comments" button
                commentsButton.addEventListener("click", () => {
                    fetchComments(post.id, commentsContainer);
                });


                postElement.appendChild(titleElement);
                postElement.appendChild(textElement);
                postElement.appendChild(userElement);
                postElement.appendChild(likesButton);
                postElement.appendChild(commentsButton);
                postElement.appendChild(commentsContainer)

                feed.appendChild(postElement);
            });
        })
        .catch((error) => {
            console.error("Error fetching posts:", error);
        });
});


function createLike(postId, userId) {
    var formData = {
        postId: postId,
        userId: userId,
    };

    var jsonData = JSON.stringify(formData);

    fetch("http://localhost:8080/posts/likes", {
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
                alert("Error creating like. Please try again.");
            }
        })
        .catch(error => {
            console.error("Error:", error);
            alert("An error occurred. Please try again later.");
        });
}