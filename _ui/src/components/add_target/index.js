import { Grid, Typography, AppBar, Toolbar, useScrollTrigger } from "@mui/material";
import { useState, cloneElement } from "react";
import Form from "./form";
import axios from 'axios'
import { ServerAddr } from '../server'
import { useDispatch, useSelector } from 'react-redux'
import { ACTION_REMOVE_TARGETS } from "../state_actions";
import { useNavigate } from "react-router-dom";

const defaultStat = {
    loading: false,
    error: ''
}
const defaultValues = {
    name: '',
    targettype: '',
    protocol: '',
    address: '',
    port: 0,
    pinginterval: '',
    pingtimeout: '',
}

function ElevationScroll(props) {
    const { children } = props;
    const trigger = useScrollTrigger({
        disableHysteresis: true,
        threshold: 0,
    });

    return cloneElement(children, {
        elevation: trigger ? 4 : 0,
    });
}

export default function AddTarget(props) {
    const [values, setValues] = useState(defaultValues)
    const [stat, setStat] = useState(defaultStat)
    const user = useSelector(state => state.user)
    const dispatch = useDispatch()
    const navigate = useNavigate()

    const changeValues = (event) => {
        setValues({
            ...values,
            [event.target.name]: event.target.value
        })
        setStat(defaultStat)
    }

    const handleFormSubmit = (event) => {
        event.preventDefault()
        setStat({
            ...defaultStat,
            loading: true,
        })
        axios.post(ServerAddr + '/target/', {
            name: values.name,
            target_type: values.targettype,
            protocol: values.protocol,
            host_address: values.address,
            targettype: values.port,
            ping_interval: values.pinginterval,
            ping_timeout: values.pingtimeout,
        }, {
            headers: {
                'Content-Type': 'application/json',
                'Authorization': user.token
            }
        })
            .then(response => {
                setStat(defaultStat)
                dispatch({ type: ACTION_REMOVE_TARGETS })
                navigate('/targets')
            })
            .catch(error => {
                setStat({
                    ...stat,
                    loading: false,
                    error: error.response.data.message
                })
            });
    }

    return <Grid
        container
        direction="column"
        justifyContent="flex-start"
        alignItems="flex-start"
    >

        <ElevationScroll {...props}>
            <AppBar
                position="fixed"
                sx={{
                    width: `calc(100% - 240px)`,
                    bgcolor: 'white',
                }}
            >
                <Toolbar>
                    <Typography
                        variant="h5"
                        color="textPrimary"
                    >
                        Add New Target
                    </Typography>
                </Toolbar>
            </AppBar>
        </ElevationScroll>
        <Toolbar />
        <Form
            values={values}
            changeValues={changeValues}
            handleFormSubmit={handleFormSubmit}
            stat={stat}
        />
    </Grid>
}