import { NextPage } from "next";
import { useEffect } from "react";
import serverConfig from "../config/server";

const Packages: NextPage = () => {
    useEffect(() => {
        fetch(`${serverConfig.host}/api/packages`)
    }, [])

    return <>packages</>
}

export default Packages