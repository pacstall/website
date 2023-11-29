import Locale from './locale'

export default <Locale>{
    home: {
        title: 'Pacstall',
        subtitle: 'AUR untuk Ubuntu',
        cards: {
            whyDifferent: {
                title: 'Mengapa Pacstall berbeda dengan semua manajer paket lain?',
                description:
                    'Pacstall menggunakan basis Ubuntu stabil, namun diperbolehkan untuk menggunakan ' +
                    'perangkat lunak mutakhir dengan sedikit bahkan tanpa kompromi, jadi ' +
                    'tidak perlu risau tentang tambalan keamanan atau fitur-fitur baru.',
            },
            howItWorks: {
                title: 'Bagaimana Pacstall bekerja?',
                description:
                    ' Pacstall menggunakan <0>pacscripts</0> (mirip' +
                    ' PKGBUILD) yang berisi hal-hal yang diperlukan untuk membuat paket,' +
                    ' dan mengubahnya menjadi program di sistem.',
            },
        },
        installationInstructions: 'Panduan Instalasi',
        showcase: {
            title: 'Pratinjau',
            packageSearch: 'Mencari Paket',
        },
    },
    navbar: {
        title: 'Pacstall',
        contribute: {
            title: 'Kontribusi',
            workOnFeatures: 'Pengerjaan fitur baru',
            helpTranslate: 'Bantuan dalam penerjemahan',
            becomeAMaintainer: 'Menjadi pemelihara paket',
        },
        social: {
            title: 'Media Sosial',
            discord: 'Discord',
            matrix: 'Matrix',
            reddit: 'Reddit',
            lemmy: 'Lemmy',
            mastodon: 'Mastodon',
        },
        browse: {
            title: 'Telusuri Paket',
        },
        privacy: {
            title: 'Kebijakan Privasi',
        },
        install: 'Pasang',
    },
    cookieConsent: {
        title: 'Catatan kuki singkatnya',
        paragraphs: [
            'Hai, iya kami menggunakan kuki',
            'Kami tidak ingin menyebabkanmu salah paham atau informasi yang membingungkan. Kami hanya menggunakan kuki untuk fitur-fitur esensial seperti setelan tema, pelokalan, dan autentikasi.',
            'Kebijakan privasi lengkap dapat dibaca <0>disini <1/></0>.',
            'Dengan lanjut menggunakan situs ini, <strong>persetujuan terhadap kebijakan privasi kami telah disetujui</strong>.',
        ],
        accept: 'Baiklah',
    },
    packageSearch: {
        dropdown: {
            package: 'Paket',
            packageTooltip: 'Mencari nama paket dan deskripsinya',
            maintainer: 'Pemelihara',
            maintainerTooltip: 'Mencari dengan nama pemelihara dan surelnya',
        },
        table: {
            name: 'Nama',
            maintainer: 'Pemelihara',
            version: 'Versi',
            install: 'Pasang',
        },
        versionTooltip: {
            notInRegistry: 'Paket ini tidak terdaftar dalam Repology',
            latest: 'Paket ini sudah versi terbaru',
            patch: 'Paket ini mendapatkan pembaruan tambalan keamanan',
            minor: 'Paket ini memiliki pembaruan minor',
            major: 'Paket ini menyediakan pembaruan mayor',
            isGit: 'Paket ini adalah paket Git',
        },
        noResults:
            'Tidak menemukan paket yang diinginkan? <0>Buat permintaan!</0>',
        search: 'Cari',
        orphaned: 'Terlantar',
        maintainerTooltip: {
            maintainedBy: 'Paket ini dipelihara oleh {{ name }}',
            noMaintainer: 'Paket ini tidak sedang dipelihara',
        },
        pagination: {
            previous: 'sebelumnya',
            next: 'selanjutnya',
        },
    },
    packageDetails: {
        table: {
            name: 'Nama',
            version: 'Versi',
            maintainer: 'Pemelihara',
            dependencies: 'Dependensi',
            requiredBy: 'Dibutuhkan oleh',
            lastUpdatedAt: 'Diperbarui pada',
        },
        orphaned: 'Ditelantarkan',
        noResults: 'Tidak ditemukan',
        openInGithub: 'Buka di Github',
        view: 'Lihat',
        howToInstall: {
            title: 'Cara pemasangan',
            step1: 'Step 1: Siapkan Pacstall',
            step2: 'Step 2: Pasang {{ name }}',
        },
        dependenciesModal: {
            title: 'Dependensi',
            buildDependencies: 'Dependensi Pembangunan',
            optionalDependencies: 'Dependensi Opsional',
            runtimeDependencies: 'Dependensi Pengeksekusi',
            pacstallDependencies: 'Dependensi Pacstall',
            name: 'Nama',
            close: 'Tutup',
            provider: 'Penyedia',
            noDescription: 'Deskripsi tidak tersedia',
            version: 'Versi',
        },
        requiredByModal: {
            title: 'Disediakan oleh',
            name: 'Nama',
            provider: 'Penyedia',
            close: 'Tutup',
            noDescription: 'Deskripsi tidak tersedia',
        },
    },
}
