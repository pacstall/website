import Locale from './locale'

export default <Locale>{
    home: {
        title: 'Pacstall',
        subtitle: 'Ubuntu için AUR',
        cards: {
            whyDifferent: {
                title: 'Bu diğer paket yöneticilerinden neden farklı?',
                description:
                    'Pacstall Ubuntu`nun stabil tabanını ve son sürüm yazılımları +
                    "hiçbir taviz vermeden kullanmanızı sağlar. Bu şekilde sizi " +
                    'güvenlik güncellemeleri yada yeni özellikler konusunda büyük' +
                    'bir endişeden kurtarır ve kolaylık sağlar.',
            },
            howItWorks: {
                title: 'O zaman bu şey nasıl çalışıyor?',
                description:
                    ' Pacstall <0>pacscripts</0> olarak bilinen (PKGBUILD dosyalarına' +
                    ' benzer) ve paketleri derlemek için gerekli olan şeyleri içeriğinde' +
                    ' barındıran ve onları kullanarak paketleri sisteminize entegre eder',
            },
        },
        installationInstructions: 'Kurulum Talimatları:',
        showcase: {
            title: 'Vitrin',
            packageSearch: 'Paketleri ara',
        },
    },
    navbar: {
        title: 'Pacstall',
        contribute: {
            title: 'Katkıda bulun',
            workOnFeatures: 'Yeni özellikler üzerinde çalış',
            helpTranslate: 'Çevirilerde bize yardımcı ol',
            becomeAMaintainer: 'Paket bakıcısı ol',
        },
        social: {
            title: 'Social Networks',
            discord: 'Discord',
            matrix: 'Matrix',
            reddit: 'Reddit',
            lemmy: 'Lemmy',
            mastodon: 'Mastodon',
        },
        browse: {
            title: 'Paketleri keşfet',
        },
        privacy: {
            title: 'Gizlilik Politikası',
        },
        install: 'İndir',
    },
    cookieConsent: {
        title: 'Çerezler TL;DR',
        paragraphs: [
            'Evet, biz bu sitede çerezleri kullanıyorum.',
            "Yanlış anlaşılmaya sebep olmak istemiyoruz. Çerezleri yalnızca tema ayarları, bölgesel bilgiler ve kimlik doğrulama için kullanıyoruz",
            'Gizlilik politikamızı <0>buradan <1/></0> okuyabilirsiniz.',
            'Bu siteyi kullanarak <strong>gizlilik politikamızı kabul etmiş olursunuz</strong>.',
        ],
        accept: 'Tamamdır',
    },
    packageSearch: {
        dropdown: {
            package: 'Paket',
            packageTooltip: 'Paket isimleri ve açıklamalarını arar',
            maintainer: 'Bakıcı',
            maintainerTooltip: 'Bakıcıların isimleri ve e-postaları ile arar',
        },
        table: {
            name: 'İsim',
            maintainer: 'Bakıcı',
            version: 'Sürüm',
            install: 'İndir',
        },
        versionTooltip: {
            notInRegistry: 'This package is not in the Repology registry',
            latest: 'Bu paket en güncel sürümde',
            patch: 'Bu paket için bir yama mevcut',
            minor: 'Bu paket için küçük bir güncelleme mevcut',
            major: 'Bu paket için büyük bir güncelleme mevcut',
            isGit: 'Bu paket bir Git paketi',
        },
        noResults: 'Aradığınızı bulamıyormusunuz? <0>Bir istek oluşturun!</0>',
        search: 'Ara',
        orphaned: 'Terk edilmiş',
        maintainerTooltip: {
            maintainedBy: 'Bu paket {{ name }} tarafından sürdürülmektedir',
            noMaintainer: 'Bu paket sürdürülmemektedir',
        },
        pagination: {
            previous: 'önceki',
            next: 'sonraki',
        },
    },
    packageDetails: {
        table: {
            name: 'İsim',
            version: 'Sürüm',
            maintainer: 'Bakıcı',
            dependencies: 'Gereksinimler',
            requiredBy: 'Taradından gerekli',
        },
        orphaned: 'Terk edilmiş',
        noResults: 'Hiçbirşey',
        openInGithub: 'Github üzerinde aç',
        view: 'İncele',
        howToInstall: {
            title: 'Nasıl kurulur',
            step1: "Adım 1: Pacstall`ı kur",
            step2: 'Adım 2: {{ name }} paketini kur',
        },
        dependenciesModal: {
            title: 'Gereksinimler',
            buildDependencies: 'Derleme Gereksinimleri',
            optionalDependencies: 'Tercihe bağlı Gereksinimler',
            runtimeDependencies: 'Çalıştırma Gereksinimleri',
            pacstallDependencies: 'Pacstall Gereksinimleri',
            name: 'İsim',
            close: 'Kapat',
            provider: 'Sağlayıcı',
            noDescription: 'Açıklama mevcut değil',
        },
        requiredByModal: {
            title: 'Tarafından Gerekli',
            name: 'İsim',
            provider: 'Sağlayıcı',
            close: 'Kapat',
            noDescription: 'Açıklama mevcut değil',
        },
    },
}
