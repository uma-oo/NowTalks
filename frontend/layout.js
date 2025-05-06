import { getPosts } from "./api/posts"

export function renderLayout() {
    let header = document.createElement('header')
    let aside = document.createElement('aside') 
    let main =  document.createElement('main') 

    document.body.append(header)
    document.body.append(aside)
    document.body.append(header)
}



export function renderForm(target, formJson, id) {
    console.log('afsdf')
    let form = document.createElement('form')
    form.id = id
    console.log(form)
    console.log(formJson)
    target.append(form)
}