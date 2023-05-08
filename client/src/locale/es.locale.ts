import Locale from './locale'

export default <Locale>{
    home: {
        title: 'Pacstall',
        subtitle: 'El AUR para Ubuntu',
        cards: {
            whyDifferent: {
                title: 'Why is this any different than any other package manager?',
                description:
                    'Pacstall uses the stable base of Ubuntu but allows you to use ' +
                    "bleeding edge software with little to no compromises, so you don't " +
                    'have to worry about security patches or new features.',
            },
            howItWorks: {
                title: '¿Cómo funciona entonces?',
                description:
                    ' Pacstall takes in files known as <link>pacscripts</link> (similar' +
                    ' to PKGBUILDs) that contain the necessary contents to build packages,' +
                    ' and builds them into executables on your system.',
            },
        },
        installationInstructions: 'Instrucciones de Instalación',
        showcase: {
            title: 'Escaparate',
            packageSearch: 'Búsqueda de Paquetes',
        },
    },
    navbar: {
        title: 'Pacstall',
        contribute: {
            title: 'Contribuir',
            workOnFeatures: 'Trabaja en nuevas características',
            helpTranslate: 'Ayuda con las traducciones',
            becomeAMaintainer: 'Conviértase en un mantenedor de paquetes',
        },
        social: {
            title: 'Redes Sociales',
            discord: 'Discord',
            matrix: 'Matrix',
            reddit: 'Reddit',
            mastodon: 'Mastodon',
        },
        browse: {
            title: 'Explorar Paquetes',
        },
        privacy: {
            title: 'La Política de Privacidad',
        },
        install: 'Instalar',
    },
}
