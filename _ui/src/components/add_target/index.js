import { Grid, Typography } from "@mui/material";
import { useState } from "react";
import Form from "./form";

const defaultStat = {
    loading: false,
    error: ''
}

export default function AddTarget() {
    const [values, setValues] = useState({
        name: '',
        targettype: '',
        protocol: '',
        pinginterval: '',
        pingtimeout: '',
    })
    const changeValues = (event) => {
        setValues({
            ...values,
            [event.target.name]: event.target.value
        })
        setStat(defaultStat)
    }
    const [stat, setStat] = useState(defaultStat)
    const handleFormSubmit = (event) => {
        event.preventDefault()
        setStat({
            ...defaultStat,
            loading: true,
        })
    }

    return <Grid
        container
        direction="column"
        justifyContent="flex-start"
        alignItems="flex-start"
    >
        <Typography
            variant="h5"
            sx={{ mb: 3 }}
        >
            Add New Target
        </Typography>
        <Form
            values={values}
            changeValues={changeValues}
            handleFormSubmit={handleFormSubmit}
            stat={stat}
        />
    </Grid>
}