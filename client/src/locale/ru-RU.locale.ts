import Locale from './locale'

export default <Locale>{
    home: {
        title: 'Паксталл',
        subtitle: 'AUR для Ubuntu',
        cards: {
            whyDifferent: {
                title: 'Чем это отличается от любого другого пакетного менеджера?',
                description:
                    'Паксталл использует стабильную базу Ubuntu, но позволяет использовать ' +
                    "новейшие программы почти что без компромиссов, так что вам не " +
                    'нужно беспокоиться о патчах безопасности или новом функционале.',
            },
            howItWorks: {
                title: 'Тогда как это работает?',
                description:
                    ' Паксталл использует файлы, известные как <0>пакскрипты</0> (похоже' +
                    ' на PKGBUILD-ы) которые содержат обязательные компоненты для сборки пакетов,' +
                    ' и собирает из них исполняемые файлы на вашей системе.',
            },
        },
        installationInstructions: 'Инструкция по установке',
        showcase: {
            title: 'Шоукейс',
            packageSearch: 'Поиск пакетов',
        },
    },
    navbar: {
        title: 'Pacstall',
        contribute: {
            title: 'Внести вклад',
            workOnFeatures: 'Работать над новым функционалом',
            helpTranslate: 'Помочь с переводом',
            becomeAMaintainer: 'Станьте сопровождающим пакета',
        },
        social: {
            title: 'Соцсети',
            discord: 'Дискорд',
            matrix: 'Матрикс',
            reddit: 'Реддит',
            lemmy: 'Лемми',
            mastodon: 'Мастодон',
        },
        browse: {
            title: 'Просмотреть пакеты',
        },
        privacy: {
            title: 'Политика конфиденциальности',
        },
        install: 'Установить',
    },
    cookieConsent: {
        title: 'Уведомление о куки',
        paragraphs: [
            "Привет, мы используем cookie-файлы на нашем сайте.",
            "Мы не любим давать вводящую в заблужение или ложную информацию. Мы используем cookie только для основных функций, таких как настройки темы, переводы и аутентификация.",
            'Вы можете прочитать полную политику конфиденциальности <0>здесь <1/></0>.',
            'Продолжая пользоваться этим сайтом, вы <strong>соглашаетесь с политикой конфиденциальности</strong>.',
        ],
        accept: 'Окей',
    },
    packageSearch: {
        dropdown: {
            package: 'Пакет',
            packageTooltip: 'Ищет по именам пакетов и описаниям',
            maintainer: 'Сопровождающий',
            maintainerTooltip: 'Ищет по именам сопровождающих и почтам',
        },
        table: {
            name: 'Название',
            maintainer: 'Сопровождающий',
            version: 'Версия',
            install: 'Установить',
        },
        versionTooltip: {
            notInRegistry: 'Этот пакет не в реестре Repology',
            latest: 'Этот пакет последней версии',
            patch: 'Для этого пакета есть доступный патч',
            minor: 'Для этого пакета есть незначительное обновление',
            major: 'Для этого пакета есть крупное обновление',
            isGit: 'Этот пакет - это Git-пакет',
        },
        noResults: 'Не находите то, что хотели? <0>Создайте запрос!</0>',
        search: 'Поиск',
        orphaned: 'Осиротевший',
        maintainerTooltip: {
            maintainedBy: 'Этот пакет обслуживается {{ name }}',
            noMaintainer: 'Этот пакет не обслуживается',
        },
        pagination: {
            previous: 'Назад',
            next: 'Вперёд',
        },
    },
    packageDetails: {
        table: {
            name: 'Название',
            version: 'Версия',
            maintainer: 'Разработчик',
            dependencies: 'Зависимости',
            requiredBy: 'Требуется',
            lastUpdatedAt: 'Последнее обновление',
        },
        orphaned: 'Осиротевший',
        noResults: 'Нет',
        openInGithub: 'Открыть на Github',
        view: 'Посмотреть',
        howToInstall: {
            title: 'Как скачать',
            step1: 'Шаг 1: Установить Паксталл',
            step2: 'Шаг 2: Установить {{ name }}',
        },
        dependenciesModal: {
            title: 'Зависимости',
            buildDependencies: 'Зависимости для сборки',
            optionalDependencies: 'Необязательные зависимости',
            runtimeDependencies: 'Зависимости среды выполнения',
            pacstallDependencies: 'Зависимости Паксталл',
            name: 'Название',
            close: 'Закрыть',
            provider: 'Провайдер',
            noDescription: 'Нет доступного описания',
            version: 'Версия',
        },
        requiredByModal: {
            title: 'Требуется',
            name: 'Название',
            provider: 'Провайдер',
            close: 'Закрыть',
            noDescription: 'Нет доступного описания',
        },
    },
}
