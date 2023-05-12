import Locale from './locale'

export default <Locale>{
    home: {
        title: 'Pacstall',
        subtitle: 'AUR dla Ubuntu',
        cards: {
            whyDifferent: {
                title: 'Dlaczego różni się to od innych menedżerów pakietów',
                description:
                    'Pacstall korzysta ze stabilnej bazy Ubuntu, ale pozwala na użycie ' +
                    'najnowszego oprogramowania z niewielkimi lub zerowymi kompromisami, więc nie musisz ' +
                    'się martwić o poprawki bezpieczeństwa lub nowe funkcje.',
            },
            howItWorks: {
                title: 'Jak to działa?',
                description:
                    ' Pacstall przyjmuje pliki znane jako <0>pacscripts</0> (podobny ' +
                    ' do PKGBUILD), które zawierają zawartość niezbędną do zbudowania pakietów' +
                    ' i buduje je jako pliki wykonywalne w twoim systemie.',
            },
        },
        installationInstructions: 'Jak zainstalować?',
        showcase: {
            title: 'Prezentacja Pacstall',
            packageSearch: 'Wyszukiwanie Pakietów',
        },
    },
    navbar: {
        title: 'Pacstall',
        contribute: {
            title: 'Bierz udział!',
            workOnFeatures: 'Pracuj nad nowymi funkcjami',
            helpTranslate: 'Pomóż z tłumaczeniem',
            becomeAMaintainer: 'Zostań twórcą pakietu',
        },
        social: {
            title: 'Media spolecznościowe',
            discord: 'Discord',
            matrix: 'Matrix',
            reddit: 'Reddit',
            mastodon: 'Mastodon',
        },
        browse: {
            title: 'Przeglądaj Pakiety',
        },
        privacy: {
            title: 'Polityka prywatności',
        },
        install: 'Zainstaluj',
    },
    cookieConsent: {
        title: 'Informacja o plikach cookie itp.',
        paragraphs: [
            'Cześć, tak, używamy plików cookie.',
            'Nie lubimy podawać wprowadzających w błąd lub mylących informacji. Używamy plików cookie tylko do niezbędnych funkcji, takich jak ustawienia motywu, lokalizacja i uwierzytelnianie.',
            'Możesz przeczytać pełną politykę prywatności <0>tutaj <1/></0>.',
            'Kontynuując korzystanie z tej witryny, <strong>wyrażasz zgodę na politykę prywatności</strong>.',
        ],
        accept: 'ok, fajnie',
    },
    packageSearch: {
        dropdown: {
            package: 'Pakiet',
            packageTooltip: 'Wyszukuje nazwy i opisy pakietów',
            maintainer: 'Twórca',
            maintainerTooltip:
                'Wyszukiwanie według nazw twórców i adresów e-mail',
        },
        table: {
            name: 'Imie',
            maintainer: 'Twórca',
            version: 'Wersja',
            install: 'Zainstaluj',
        },
        versionTooltip: {
            notInRegistry: 'Ten pakiet nie znajduje się w rejestrze Repology',
            latest: 'Ten pakiet jest najnowszą wersją',
            patch: 'Ten pakiet ma dostępną aktualizację poprawki',
            minor: 'Ten pakiet ma dostępną niewielką aktualizację',
            major: 'Ten pakiet ma dostępną główną aktualizację',
            isGit: 'Ten pakiet jest pakietem Git',
        },
        noResults:
            'Nie możesz znaleźć tego, czego szukasz? <0>Utwórz prośbę!</0>',
        search: 'Szukaj',
        orphaned: 'Osierocony',
        maintainerTooltip: {
            maintainedBy: 'Ten pakiet jest utrzymywany przez {{ name }}',
            noMaintainer: 'Ten pakiet nie jest utrzymywany',
        },
        pagination: {
            previous: 'poprzednia',
            next: 'następna',
        },
    },
    packageDetails: {
        table: {
            name: 'Nazwa',
            version: 'Wersja',
            maintainer: 'Twórca',
            dependencies: 'Zależności',
            requiredBy: 'Wymagany przez',
        },
        orphaned: 'Nie utrzymywany',
        noResults: 'Brak',
        openInGithub: 'Otwórz w GitHubie',
        view: 'Zobacz',
        howToInstall: {
            title: 'Jak zainstalować?',
            step1: 'Krok 1: Konfiguracja Pacstall',
            step2: 'Krok 2: Zainstaluj {{ name }}',
        },
        dependenciesModal: {
            title: 'Zależności',
            buildDependencies: 'Zależności budowania',
            optionalDependencies: 'Opcjonalne zależności',
            runtimeDependencies: 'Zależności w czasie rzeczywistym',
            pacstallDependencies: 'Zależności Pacstall',
            name: 'Nazwa',
            close: 'Zamknij',
            provider: 'Dostawca',
            noDescription: 'Brak opisu',
        },
        requiredByModal: {
            title: 'Wymagany przez',
            name: 'Nazwa',
            provider: 'Dostawca',
            close: 'Zamknij',
            noDescription: 'Brak opisu',
        },
    },
}
