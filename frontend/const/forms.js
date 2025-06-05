export const registerFom = {
    elements: [
        {
            tag: 'input',
            label: 'NickName:',
            attributes: {
                required: 'true',
                type: 'text',
                id: 'nickname',
                name: 'nickname',
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
                width: '100%'
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
                width: '100%'
            }
        },
        {
            tag: 'input',
            label: 'FirstName',
            attributes: {
                required: 'true',
                type: 'text',
                id: 'firstname',
                name: 'firstname',
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
                id: 'lastname',
                name: 'lastname',
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
                id: 'password',
                name: 'password',
                pattern: "^(?=.*[0-9])(?=.*[!@#$%^&*])[a-zA-Z0-9!@#$%^&*]{8,}$",
                placeholder: "Enter Your Password...",
            },
            style: {
                width: '100%'
            }
        },
        {
            tag: 'input',
            label: 'Verify Password:',
            attributes: {
                required: 'true',
                type: 'password',
                id: 'password2',
                name: 'password2',
                pattern: "^(?=.*[0-9])(?=.*[!@#$%^&*])[a-zA-Z0-9!@#$%^&*]{8,}$",
                placeholder: "Repeat Your Password...",
            },
            style: {
                width: '100%'
            }
        },
    ],
    buttons: [
        {
            type: 'submit',
            content: {
                text: "Register"
            },
            style: 'primary-btn'
        },
        {
            type: 'reset',
            content: {
                text: "Reset"
            },
            style: 'secondary-btn'
        }
    ]
}



export const loginForm = {
    elements: [
        {
            tag: 'input',
            label: 'NickName Or Email:',
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
                id: 'password',
                name: 'password',
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
            content: {
                text : 'Log In',
            },
            style: 'primary-btn'
        },
        {
            type: 'reset',
            content: {
                text : "Reset",
            },
            style: 'secondary-btn'
        }
    ]
}


export const PostForm = {
    elements: [
        {
            tag: 'input',
            // label: 'Post Title',
            attributes: {
                required: 'true',
                type: 'text',
                id: 'title',
                name: 'title',
                placeholder: "Your post title here",
            },
            style: {
                width: '100%'
            }
        },
        {
            tag: 'textarea',
            // label: 'Post Content:',
            attributes: {
                required: 'true',
                id: 'content',
                name: 'content',
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
            content: {
                text : "share",
                icon : "share"
            },
            style: ['primary-btn','row-reverse']
        },
        {
            type: 'reset',
            content: {
                text : "cancell"
            },
            style: 'secondary-btn'
        }
    ]
}

export const CommentForm = {
    elements: [
        {
            tag: 'input',
            attributes: {
                required: 'true',
                type: 'text',
                id: 'content',
                name: 'content',
                placeholder: "Add a comment...",
            },
            style: {
                // width: '100%'
            }
        },
    ],
    buttons: [
        {
            type: 'submit',
            content: {
                icon : "send"
            },
            style: 'primary-btn'
        }
    ]
}

export const MessageForm = {
    elements: [
        {
            tag: 'textarea',
            attributes: {
                required: 'true',
                type: 'text',
                id: 'chatMessage',
                name: 'chatMessage',
                placeholder: "Send a message..."
            },
            style: {
                // width: '100%'
            }
        },
    ],
    buttons: [
        {
            type: 'submit',
            content: {
                icon : "share"
            },
            style: 'primary-btn'
        }
    ]
}