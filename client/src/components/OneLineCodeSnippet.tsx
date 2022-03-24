import { chakra, SkeletonText, useColorModeValue } from "@chakra-ui/react";
import { CSSProperties, FC, MutableRefObject, useRef } from "react";
import { useRecoilState } from "recoil";
import useNotification from "../hooks/useNotification";
import { featureFlagsState } from "../state/feature-flags";
import Notification from "../types/notifications";


const onCommandCopy = async (ref: MutableRefObject<HTMLInputElement | undefined>, notify: (notification: Notification) => any) => {
    if (!ref.current) {
        return
    }

    const notifyCopied = () => notify({
        title: 'Copied to Clipboard!',
        text: 'You can now paste this command in the terminal.',
        type: 'info'
    })

    const notifyError = () => notify({
        title: 'Failed to copy to Clipboard',
        text: 'Make sure you are running an up-to-date browser.',
        type: 'warning'
    })

    ref.current.focus()
    ref.current.select()
    if (document.execCommand('copy', true)) {
        notifyCopied()
    } else if ('clipboard' in navigator) {
        try {
            await navigator.clipboard.writeText(ref.current.value);
            notifyCopied()
        } catch {
            notifyError()
        }
    } else {
        notifyError()
    }
}

const nonSelectableStyle: CSSProperties = {
    WebkitUserSelect: 'none',
    MozUserSelect: 'none',
    msUserSelect: 'none',
    userSelect: 'none'
}

const OneLineCodeSnippet: FC<{ size?: 'xs' | 'sm' | 'md' | 'lg' }> = ({ children, size }) => {
    const notify = useNotification()
    const ref = useRef<HTMLInputElement>()
    const text = children!.toString() as string
    return (
        <>
            <chakra.span
                p='1'
                px='2'
                fontFamily='JetBrains Mono'
                bg={useColorModeValue('gray.100', 'gray.900')}
                color={useColorModeValue('teal.500', 'teal.400')}
                fontWeight='500'
                onClick={() => onCommandCopy(ref, notify)}
                cursor='pointer'
                wordBreak='break-all'
                fontSize={size}
                borderRadius='md'>
                <chakra.span color='pink.400' fontWeight='800' style={nonSelectableStyle}>$ </chakra.span>{text}
            </chakra.span>
            <input ref={ref as any} type="text" style={{ display: "none", maxWidth: "1px" }} readOnly defaultValue={text} />
        </>
    )
}

export default OneLineCodeSnippet

export const SmartCodeSnippetInstall: FC<{ size?: 'xs' | 'sm' | 'md' | 'lg', name: string }> = ({ size, name }) => {
    const [featureFlags] = useRecoilState(featureFlagsState)
    if (featureFlags.loading) {
        return <SkeletonText noOfLines={1} />
    }

    const code = () => {
        if (featureFlags.flags?.oldSyntax) {
            return `pacstall -I ${name}`
        }

        return `sudo pacstall install ${name}`
    }

    return <OneLineCodeSnippet size={size}>{code()}</OneLineCodeSnippet>
}