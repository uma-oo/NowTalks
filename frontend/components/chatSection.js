export function createChatSections(){
    let chatSection = document.createElement('section')
    chatSection.classList.add("chat_section","tab_section")
    chatSection.append("Chat Section")
    return chatSection;
}