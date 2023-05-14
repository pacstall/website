import Locale from './locale'

export default <Locale>{
    home: {
        title: 'पैकस्टॉल"',
        subtitle: 'उबंटू के लिए मानक एउआर (AUR)',
        cards: {
            whyDifferent: {
                title: 'यह अन्य पैकेज प्रबंधकों से किस प्रकार भिन्न है?',
                description:
                    'पैकस्टॉल आपको उबंटू की स्थिरता से समझौता किए बिना ब्लीडिंग एज सॉफ़्टवेयर का उपयोग करने की क्षमता ' +
                    "देता है, इसलिए आपको सॉफ़्टवेयर सुरक्षा पैच या नई फीचर के बारे में चिंता करने की आवश्यकता नहीं है।"
            },
            howItWorks: {
                title: 'तो यह कैसे काम करता है?',
                description:
                    'पैकस्टॉल <0>pacscripts</0> (जैसे PKGBUILDs) नामक फाइलों का उपयोग पैकेज बनाने के लिए ' +
                    "आवश्यक सामग्री होती है और उससे आपके सिस्टम पर एक एक्सीक्यूटेबल बनाता है।"
            },
        },
        installationInstructions: 'अधिष्ठापित करने का निर्देश',
        showcase: {
            title: 'प्रदर्शनी',
            packageSearch: 'पैकेज खोज',
        },
    },
    navbar: {
        title: 'पैकस्टॉल',
        contribute: {
            title: 'योगदान करें',
            workOnFeatures: 'नई सुविधाएं बनाने में सहायता करें',
            helpTranslate: 'इस वेबसाइट का अनुवाद करने में सहायता करें',
            becomeAMaintainer: 'एक पैकेज रखरखावकर्ता बनें',
        },
        social: {
            title: 'सामाजिक संपर्क',
            discord: 'डिस्कॉर्ड',
            matrix: 'मैट्रिक्स',
            reddit: 'रेड्डिट',
            mastodon: 'सटोडॉन',
        },
        browse: {
            title: 'पैकेजो को ब्राउज़ करें',
        },
        privacy: {
            title: 'गोपनीयता नीति',
        },
        install: 'अधिष्ठापित करे',
    },
    cookieConsent: {
        title: 'कुकी सूचना',
        paragraphs: [
            'नमस्ते, हाँ, हम कुकीज़ का उपयोग करते हैं।',
            "हम आपको भ्रमित करने वाली जानकारी देना पसंद नहीं करते हैं। हम केवल थीम सेटिंग्स, स्थानीयकरण और प्रमाणीकरण जैसी आवश्यक सुविधाओं के लिए कुकीज़ का उपयोग करते हैं।",
            'आप पूरी गोपनीयता नीति <0>यहां<1/></0> पढ़ सकते हैं।',
            'इस साइट का उपयोग जारी रखकर आप <strong>गोपनीयता नीति को अपनी सहमति देते हैं</strong>।',
        ],
        accept: 'ठीक है',
    },
    packageSearch: {
        dropdown: {
            package: 'पैकेज',
            packageTooltip: 'पैकेजो के नाम और विवरण में खोजता है',
            maintainer: 'ररखरखावकर्ता',
            maintainerTooltip: 'खरखाव करता के नाम और ईमेल द्वारा खोजता है',
        },
        table: {
            name: 'नाम',
            maintainer: 'खरखाव करता',
            version: 'संस्करण',
            install: 'अधिष्ठापित करनाे का आदेश',
        },
        versionTooltip: {
            notInRegistry: 'यह पैकेज रेपोलॉजी रजिस्ट्री में नहीं है',
            latest: 'यह पैकेज नवीनतम संस्करण में है',
            patch: 'इस पैकेज में एक पैच अद्यतन लंबित है',
            minor: 'इस पैकेज में एक मामूली अद्यतन लंबित है',
            major: 'इस पैकेज में एक बड़ा अद्यतन लंबित है',
            isGit: 'यह पैकेज एक गिट पैकेज है',
        },
        noResults: 'आप जो खोज रहे हैं वह नहीं मिल रहा है? <0>अनुरोध करें!</0>',
        search: 'खोजें',
        orphaned: 'ऑर्फंड (कोई रखरखावकर्ता नहीं है)',
        maintainerTooltip: {
            maintainedBy: 'इस पैकेज का रखरखावकर्ता {{नाम}} है',
            noMaintainer: 'इस पैकेज का रखरखाव नहीं किया जा रहा है',
        },
        pagination: {
            previous: 'पिछला',
            next: 'अगला',
        },
    },
    packageDetails: {
        table: {
            name: 'नाम',
            version: 'संस्करण',
            maintainer: 'रखरखावकर्ता',
            dependencies: 'निर्भरताएँ',
            requiredBy: 'द्वारा अपेक्षित',
        },
        orphaned: 'ऑर्फंड (कोई रखरखावकर्ता नहीं है)',
        noResults: 'शून्य',
        openInGithub: 'गिटहब में खोलें',
        view: 'देखो',
        howToInstall: {
            title: 'अधिष्ठापित करने का निर्देश',
            step1: 'चरण १: पैकस्टॉल अधिष्ठापित करे',
            step2: 'चरण २: {{ name }} अधिष्ठापित करे',
        },
        dependenciesModal: {
            title: 'निर्भरताएँ',
            buildDependencies: 'निर्माण के दौरान निर्भरताएँ',
            optionalDependencies: 'वैकल्पिक निर्भरताएँ',
            runtimeDependencies: 'निष्पादन समय के निर्भरताएँ',
            pacstallDependencies: 'पैकस्टॉल निर्भरताएँ',
            name: 'नाम',
            close: 'बंद करे',
            provider: 'प्रदाता',
            noDescription: 'कोई विवरण उपलब्ध नहीं है',
        },
        requiredByModal: {
            title: 'द्वारा अपेक्षित',
            name: 'नाम',
            provider: 'प्रदाता',
            close: 'बंद करे',
            noDescription: 'कोई विवरण उपलब्ध नहीं है',
        },
    },
}
