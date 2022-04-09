const untypedWindow: any = window

// Opera 8.0+
const isOpera =
    (!!untypedWindow.opr && !!untypedWindow.opr.addons) ||
    !!untypedWindow.opera ||
    navigator.userAgent.indexOf(' OPR/') >= 0

// Firefox 1.0+
const isFirefox = typeof untypedWindow.InstallTrigger !== 'undefined'

// Safari 3.0+ "[object HTMLElementConstructor]"
const isSafari =
    /constructor/i.test(untypedWindow.HTMLElement) ||
    (function (p) {
        return p.toString() === '[object SafariRemoteNotification]'
    })(
        !untypedWindow['safari'] ||
            (typeof untypedWindow.safari !== 'undefined' &&
                untypedWindow['safari'].pushNotification),
    )

// Internet Explorer 6-11
const isIE = /*@cc_on!@*/ false || !!untypedWindow.document.documentMode

// Edge 20+
const isEdge = !isIE && !!untypedWindow.StyleMedia

// Chrome 1 - 79
const isChrome =
    !!untypedWindow.chrome &&
    (!!untypedWindow.chrome.webstore || !!untypedWindow.chrome.runtime)

// Edge (based on chromium) detection
const isEdgeChromium = isChrome && navigator.userAgent.indexOf('Edg') != -1

// Blink engine detection
const isBlink = (isChrome || isOpera) && !!untypedWindow.CSS

const browser = {
    isBlink,
    isChrome,
    isEdge,
    isEdgeChromium,
    isFirefox,
    isIE,
    isOpera,
    isSafari,
} as const

export default browser
