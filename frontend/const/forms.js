export const registerFom = {
    elements: [
        {
            tag: 'input',
            label: 'NickName:',
            attributes: {
                required: 'true',
                type: 'text',
                id: 'nickName',
                name: 'nickName',
                placeholder: "nickname...",
                minlength: 4,
                maxlenght: 20,
            },
            style: {
                width: '100%'
            }
        },
        {
            tag: 'select',
            label: 'Gender:',
            attributes: {
                required: 'true',
                id: 'gender',
                name: 'gender',
            },
            options: ["Male", "Female"],
            style: {
                width: '50%'
            }
        },
        {
            tag: 'input',
            label: 'Age:',
            attributes: {
                required: 'true',
                type: 'number',
                id: 'age',
                name: 'age',
                min: 18,
                default: 18
            },
            style: {
                width: '50%'
            }
        },
        {
            tag: 'input',
            label: 'FirstName',
            attributes: {
                required: 'true',
                type: 'text',
                id: 'firstName',
                name: 'firstName',
                placeholder: "Enter Your First Name...",
            },
            style: {
                width: '100%'
            }
        },
        {
            tag: 'input',
            label: 'LastName',
            attributes: {
                required: 'true',
                type: 'text',
                id: 'lastName',
                name: 'lastName',
                placeholder: "Enter Your Last Name...",
            },
            style: {
                width: '100%'
            }
        },
        {
            tag: 'input',
            label: 'Email',
            attributes: {
                required: 'true',
                type: 'email',
                id: 'email',
                name: 'email',
                pattern: "^[a-zA-Z0-9._%+\\-]+@[a-zA-Z0-9.\\-]+\\.[a-zA-Z]{2,}$",
                placeholder: "Enter Your Email...",
            },
            style: {
                width: '100%'
            }
        },
        {
            tag: 'input',
            label: 'Password:',
            attributes: {
                required: 'true',
                type: 'password',
                id: 'userPassword',
                name: 'userPassword',
                pattern: "^(?=.*[0-9])(?=.*[!@#$%^&*])[a-zA-Z0-9!@#$%^&*]{6,16}$",
                placeholder: "Enter Your Password...",
            },
            style: {
                width: '100%'
            }
        },
    ],
    buttons: [
        {
            type: 'submit',
            content: 'Register',
            style: 'primary-btn'
        },
        {
            type: 'reset',
            content: 'Cancel',
            style: 'secondary-btn'
        }
    ]
}



export const loginForm = {
    elements: [
        {
            tag: 'input',
            label: 'NickName/Email',
            attributes: {
                required: 'true',
                type: 'text',
                id: 'login',
                name: 'login',
                placeholder: "Enter NickName or Email...",
            },
            style: {
                width: '100%'
            }
        },
        {
            tag: 'input',
            label: 'Password:',
            attributes: {
                required: 'true',
                type: 'password',
                id: 'userPassword',
                name: 'userPassword',
                placeholder: "Enter Your Password...",
            },
            style: {
                width: '100%'
            }
        },
    ],
    buttons: [
        {
            type: 'submit',
            content: 'Log In',
            style: 'primary-btn'
        },
        {
            type: 'reset',
            content: 'Cancel',
            style: 'secondary-btn'
        }
    ]
}


export const PostForm = {
    elements: [
        {
            tag: 'input',
            label: 'Post Title',
            attributes: {
                required: 'true',
                type: 'text',
                id: 'postTitle',
                name: 'postTitle',
                placeholder: "Your post title here",
            },
            style: {
                width: '100%'
            }
        },
        {
            tag: 'textarea',
            label: 'Post Content:',
            attributes: {
                required: 'true',
                id: 'postContent',
                name: 'postContent',
                placeholder: "Your post content goes here...",
            },
            style: {
                width: '100%'
            }
        },
    ],
    buttons: [
        {
            type: 'submit',
            content: 'Share Post',
            style: 'primary-btn'
        },
        {
            type: 'reset',
            content: 'Cancel',
            style: 'secondary-￼btn'
        }
    ]
}

export const CommentForm = {
    elements: [
        {
            tag: 'input',
            label: 'Comment',
            attributes: {
                required: 'true',
                type: 'text',
                id: 'postComment',
                name: 'postComment',
                placeholder: "Add a comment...",
            },
            style: {
                width: '100%'
            }
        },
    ],
    buttons: [
        {
            type: 'submit',
            content: 'Send Comment',
            style: 'primary-btn'
        },
        {
            type: 'reset',
            content: 'Del',
            style: 'secondary-btn'
        }
    ]
}