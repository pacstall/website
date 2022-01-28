import { useRecoilState } from "recoil"
import Notification, { notificationsState } from "../state/notifications"


const useNotification = () => {
    const [notifications, setNotifications] = useRecoilState(notificationsState)

    return (notification: Notification) => {
        notification.id = Math.random().toString().substring(2)
        setNotifications([...notifications, notification])
        setTimeout(() => {
            setNotifications(notifications.filter(n => n !== notification))
        }, 3000)
    }
}

export default useNotification