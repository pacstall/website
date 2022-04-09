import { ToastId, useToast } from '@chakra-ui/react'
import Notification from '../types/notifications'

const duration = 5000
let toasts: ToastId[] = []
const toastLimit = 5

const useNotification = () => {
    const toast = useToast({
        position: 'top',
    })
    return (notification: Notification) => {
        ;(async () => {
            while (
                toasts.filter(id => toast.isActive(id)).length >= toastLimit
            ) {
                await new Promise(resolve => setTimeout(resolve, 100))
            }

            toasts = toasts.filter(id => toast.isActive(id))

            const id = toast({
                status: notification.type,
                title: notification.title,
                description: notification.text,
                variant: 'solid',
                duration,
            })

            toasts.push(id)
        })()
    }
}

export default useNotification
