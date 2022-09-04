import { Grid } from "@mui/material";
import { useEffect, useState } from "react";
import NonSuccess from "./non_success";
import ButtonBar from "./button_bar";
import { useSelector } from 'react-redux'
import ListTargets from "./list_targets";

export default function Targets() {
    const [state, setState] = useState({
        loading: false,
        error: "something went wrong"
    })
    const targets = useSelector(state => state.targets)

    const renderTargets = () => {
        if (state.loading || state.error !== "") {
            return <NonSuccess state={state} />
        }
        return <ListTargets targets={targets} />
    }

    const fetchTargets = () => {

    }

    useEffect(() => {
        if (targets === null || targets === undefined || targets.length === 0) {
            fetchTargets()
        }
        // need to be run only once
    }, [])

    return <>
        <Grid
            container
            direction="row"
            justifyContent="flex-start"
            alignItems="center"
        >
            <ButtonBar
                fetchTargets={fetchTargets}
            />
            {renderTargets()}
        </Grid>
    </>
}