// Функция создает боковую панель навигации
// logoText - текст возле логотипа микросервиса
// logoIconName - наименование логотипа из библиотеки Google Icons
// sidebarLinks - список ссылок для боковой панели навигации
function createSideNav(logoText, logoIconName, sidebarLinks) {
    // создание списка
    const sidebar = document.createElement('ul')
    // добавление класса sidenav из Materialize CSS
    sidebar.classList.add('sidenav')
    // назначение уникального идентификатора
    sidebar.id = "sidenav-left"
    sidebar.style = "transform: translateX(0px);"

    // Создание контейнера для логотипа боковой панели навигации
    const logoContainer = document.createElement('li')
    logoContainer.classList.add('logo-container')
    // Текст возле логотипа
    logoContainer.textContent = logoText

    // Создание логотипа
    const logoIcon = document.createElement('i')
    logoIcon.classList.add('material-icons', 'left')
    logoIcon.textContent = logoIconName

    // Добавление всех элементов в тэг 'header'
    logoContainer.appendChild(logoIcon)
    sidebar.appendChild(logoContainer)
    document.querySelector('header').appendChild(sidebar)

    sidebarLinks.forEach(link => {
        const navItem = document.createElement('li')
        navItem.classList.add('bold')
        if (window.location.pathname == link.href) {
            navItem.classList.add('active')
        }

        const itemLink = document.createElement('a')
        itemLink.classList.add('waves-effect')
        itemLink.href = link.href
        itemLink.textContent = link.label

        navItem.appendChild(itemLink)
        sidebar.appendChild(navItem)
    })
}
