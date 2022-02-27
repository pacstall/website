import { useToast } from "@chakra-ui/react"
import Notification from "../types/notifications"


const useNotification = () => {
    const toast = useToast({
        position: 'top'
    })
    return (notification: Notification) => {
        toast({
            status: notification.type,
            title: notification.title,
            description: notification.text
        })
    }
}

export default useNotification