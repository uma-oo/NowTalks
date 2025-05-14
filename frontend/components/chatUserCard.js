export function createChatUserCard(user) {
    let chatUserCard = document.createElement('div');
    chatUserCard.classList.add("chat-user-card")

    let userName = document.createElement('p')
    userName.append(user.name)
    let onlineIndicator = document.createElement('div')
    onlineIndicator.classList.add('online-indicator')
    user.online ? onlineIndicator.classList.add('online') : onlineIndicator.classList.add('offline')
    chatUserCard.append(onlineIndicator, userName)
    return chatUserCard
}