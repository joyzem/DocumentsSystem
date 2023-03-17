// Классы, используемые в этом коде, взяты из библиотеки Materialize CSS

// Создание элемента навигационной панели (navbar)
const navbar = document.createElement('nav')
// Добавление класса z-depth-0, который устанавливает размер тени, равный нулю
navbar.classList.add('z-depth-0')

// Создание обертки для навигационной панели (navWrapper)
const navWrapper = document.createElement('div')
// Добавление обертки для панели навигации, цвет - blue
navWrapper.classList.add('nav-wrapper', 'blue')
// Добавление внутреннего элемента в navbar
navbar.appendChild(navWrapper)

// Создание списка навигации (navList).  Элемент ul - неупорядоченный список
const navList = document.createElement('ul')
// Выравнивание по левому краю
navList.classList.add('left')
navWrapper.appendChild(navList)

// Создание списка ссылок навигации (navLinks)
const navLinks = [
    {
        label: 'Документы',
        href: '/documents/proxies',
    },
    {
        label: 'Товары',
        href: '/product/products',
    },
    {
        label: 'Сотрудники',
        href: '/employee/employees',
    },
    {
        label: 'Контрагенты',
        href: '/customer/customers',
    },
    {
        label: 'Организации',
        href: '/organization/organizations',
    },
    {
        label: 'Счета',
        href: '/account/accounts',
    },
];

// Добавление ссылок в список навигации (navList)
navLinks.forEach(link => {
    // Добавление элемента в неупорядоченный список navList
    const navItem = document.createElement('li')
    navItem.classList.add('nav-item')
    navList.appendChild(navItem)

    // Создание ссылки
    const navLink = document.createElement('a')
    // Текст ссылки
    navLink.textContent = link.label
    // Адрес ссылки
    navLink.href = link.href
    // Добавление ссылки в navItem
    navItem.appendChild(navLink)
});

// Создание обертки для навигационной панели (navbarContainer)
const navbarContainer = document.createElement('div')
navbarContainer.classList.add('navbar-fixed')
navbarContainer.appendChild(navbar)

// Добавление навигационной панели (navbar) на страницу
document.querySelector('header').appendChild(navbarContainer)

// Получение текущего пути
const currentPath = window.location.pathname

// Получение элементов списка навигации (navItems)
const navItems = document.querySelectorAll('.nav-item');

// Подсветка текущего пункта навигации
for (let i = 0; i < navItems.length; i++) {
    item = navItems[i]
    let tabLink = item.querySelector('a').getAttribute('href')
    rootUrl = tabLink.slice(0, tabLink.slice(1, tabLink.length).indexOf("/"))
    if (currentPath.includes(rootUrl)) {
        item.classList.add('active')
    } else {
        item.classList.remove('active')
    }
}
