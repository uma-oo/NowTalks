export const registerFom = [
    {
        tag: 'input',
        attributes: {
            type: 'text',
            id: 'nickName',
            placeholder: "nickname...",
            minlength: 4,
            maxlenght: 20,
        },
        style: {
            width: '100%'
        }
    },
    {
        tag: 'input',
        attributes: {
            type: 'number',
            id: 'age',
            min: 18,
        },
        style: {
            width: '50%'
        }
    },
    {
        tag: 'select',
        attributes: {
            id: 'gender',
        },
        options: ["Male", "Female"],
        style: {
            width: '50%'
        }
    },
    {
        tag: 'input',
        attributes: {
            type: 'text',
            id: 'firstName',
            placeholder: "Enter Your First Name...",
        },
        style: {
            width: '100%'
        }
    },
    {
        tag: 'input',
        attributes: {
            type: 'text',
            id: 'lastName',
            placeholder: "Enter Your Last Name...",
        },
        style: {
            width: '100%'
        }
    },
    {
        tag: 'input',
        type: 'email',
        id: 'email',
        props: {
            placeholder: "Enter Your Email...",
        },
        style: {
            width: '100%'
        }
    },
    {  
        tag: 'input',
        type: 'password',
        id: 'email',
        props: {
            placeholder: "Enter Your Password...",
        },
        style: {
            width: '100%'
        }
    },

]