import { FC, MutableRefObject, useReducer, useRef } from "react";
import useNotification from "../hooks/useNotification";
import { useFeatureFlags } from "../state/feature-flags";
import Notification from "../state/notifications";

const onCommandCopy = (ref: MutableRefObject<HTMLInputElement | undefined>, notify: (notification: Notification) => any) => {
    if (!ref.current) {
        return
    }

    ref.current.style.display = "inline"
    ref.current.focus()
    ref.current.select()
    if (document.execCommand('copy', true)) {
        notify({
            title: 'Copied to Clipboard!',
            text: 'You can now paste this command in the terminal.',
            type: 'success'
        })
    }
    ref.current.style.display = "none"
}

const OneLineCodeSnippet: FC<{ size?: 'xs' | 'sm' | 'md' | 'lg' }> = ({ children, size }) => {
    const notify = useNotification()
    const ref = useRef<HTMLInputElement>()
    const text = children!.toString() as string
    return (
        <>
            <span className={`px-2 ${size && size !== 'xs' && 'py-1' || ''} inline-flex text-${size || 'xs'} leading-5 rounded bg-gray-600 text-white`} style={{ cursor: "copy", display: 'block' }} onClick={() => onCommandCopy(ref, notify)}>
                <span className="text-green-300 font-semibold">$</span>
                <span className="px-2 whitespace-nowrap">{text}</span>
            </span>
            <input ref={ref as any} type="text" style={{ display: "none", maxWidth: "1px" }} readOnly defaultValue={text} />
        </>
    )
}

export default OneLineCodeSnippet

export const SmartCodeSnippetInstall: FC<{ size?: 'xs' | 'sm' | 'md' | 'lg', name: string }> = ({ size, name }) => {
    const featureFlags = useFeatureFlags()
    return <OneLineCodeSnippet size={size} >{featureFlags.flags?.oldSyntax ? `pacstall -I ${name}` : `sudo pacstall install ${name}`}</OneLineCodeSnippet>
}