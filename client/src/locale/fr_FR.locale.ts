import Locale from './locale'

export default <Locale>{
    home: {
        title: 'Pacstall',
        subtitle: 'L\'AUR pour Ubuntu',
        cards: {
            whyDifferent: {
                title: 'Pourquoi est-ce différent de tout autre gestionnaire de paquets ?',
                description:
                    'Pacstall utilise la base stable d\'Ubuntu mais vous permet d\'utiliser ' +
                    'des logiciels de pointe sans compromis, vous n\'aurez donc pas à ' +
                    'vous soucier des correctifs de sécurité ou des nouvelles fonctionnalités.',
            },
            howItWorks: {
                title: 'Comment ça fonctionne alors ?',
                description:
                    ' Pacstall prend en charge les fichiers appelés <0>pacscripts</0> (similaires ' +
                    'aux PKGBUILDs) qui contiennent le contenu nécessaire pour construire des paquets, ' +
                    'et les construit en exécutables sur votre système.',
            },
        },
        installationInstructions: 'Instructions d\'installation',
        showcase: {
            title: 'Vitrine',
            packageSearch: 'Recherche de paquets',
        },
    },
    navbar: {
        title: 'Pacstall',
        contribute: {
            title: 'Contribuer',
            workOnFeatures: 'Travailler sur de nouvelles fonctionnalités',
            helpTranslate: 'Aider avec les traductions',
            becomeAMaintainer: 'Devenir un mainteneur de paquets',
        },
        social: {
            title: 'Réseaux Sociaux',
            discord: 'Discord',
            matrix: 'Matrix',
            reddit: 'Reddit',
            mastodon: 'Mastodon',
        },
        browse: {
            title: 'Parcourir les paquets',
        },
        privacy: {
            title: 'Politique de confidentialité',
        },
        install: 'Installer',
    },
    cookieConsent: {
        title: 'Avis sur les cookies TL;DR',
        paragraphs: [
            'Bonjour, oui, nous utilisons des cookies.',
            "Nous n'aimons pas vous donner des informations trompeuses ou confuses. Nous n'utilisons des cookies que pour des fonctionnalités essentielles telles que les paramètres de thème, la localisation et l'authentification.",
            'Vous pouvez lire la politique de confidentialité complète <0>ici <1/></0>.',
            'En continuant à utiliser ce site, vous <strong>acceptez la politique de confidentialité</strong>.',
        ],
        accept: 'ok, sympa',
    },
    packageSearch: {
        dropdown: {
            package: 'Paquet',
            packageTooltip: 'Recherche dans les noms et les descriptions des paquets',
            maintainer: 'Mainteneur',
            maintainerTooltip: 'Recherche par noms et e-mails de mainteneur',
        },
        table: {
            name: 'Nom',
            maintainer: 'Mainteneur',
            version: 'Version',
            install: 'Installer',
        },
        versionTooltip: {
            notInRegistry: 'Ce paquet n\'est pas dans le registre Repology',
            latest: 'Ce paquet est la dernière version',
            patch: 'Ce paquet a une mise à jour de correctif disponible',
            minor: 'Ce paquet a une mise à jour mineure disponible',
            major: 'Ce paquet a une mise à jour majeure disponible',
            isGit: 'Ce paquet est un package Git',
        },
        noResults: 'Vous ne trouvez pas ce que vous cherchez ? <0>Créez une demande !</0>',
        search: 'Recherche',
        orphaned: 'Orphelin',
        maintainerTooltip: {
            maintainedBy: 'Ce paquet est maintenu par {{ name }}',
            noMaintainer: 'Ce paquet n\'est pas maintenu',
        },
        pagination: {
            previous: 'précédent',
            next: 'suivant',
        },
    },
    packageDetails: {
        table: {
            name: 'Nom',
            version: 'Version',
            maintainer: 'Mainteneur',
            dependencies: 'Dépendances',
            requiredBy: 'Requis par',
        },
        orphaned: 'Orphelin',
        noResults: 'Aucun',
        openInGithub: 'Ouvrir dans Github',
        view: 'Voir',
        howToInstall: {
            title: 'Comment installer',
            step1: 'Étape 1 : Configuration de Pacstall',
            step2: 'Étape 2 : Installation de {{ name }}',
        },
        dependenciesModal: {
            title: 'Dépendances',
            buildDependencies: 'Dépendances de construction',
            optionalDependencies: 'Dépendances facultatives',
            runtimeDependencies: 'Dépendances d\'exécution',
            pacstallDependencies: 'Dépendances Pacstall',
            name: 'Nom',
            close: 'Fermer',
            provider: 'Fournisseur',
            noDescription: 'Aucune description disponible',
        },
        requiredByModal: {
            title: 'Requis par',
            name: 'Nom',
            provider: 'Fournisseur',
            close: 'Fermer',
            noDescription: 'Aucune description disponible',
        },
    },
}
