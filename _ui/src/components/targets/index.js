import { Button, CircularProgress, Grid, Typography } from "@mui/material";
import { AddLink } from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { useState } from "react";

export default function Targets() {
    const navigate = useNavigate()
    const [state, setState] = useState({
        loading: false,
        error: "something went wrong"
    })

    const renderNonSuccess = (content) => {
        return <Grid
            container
            direction="column"
            justifyContent="flex-start"
            alignItems="center"
        >
            <br />
            {content}
            <br />
        </Grid>
    }

    const renderTargets = () => {
        if (state.loading) {
            return renderNonSuccess(<CircularProgress size={25} />)
        } else if (state.error !== "") {
            return renderNonSuccess(
                <Typography variant="body1" color="error">
                    {state.error}
                </Typography>
            )
        }
        return <></>
    }

    return <>
        <Grid
            container
            direction="row"
            justifyContent="flex-start"
            alignItems="center"
        >
            <Button
                variant="contained"
                startIcon={<AddLink />}
                onClick={() => navigate('new')}
            >
                Add Target
            </Button>
            {renderTargets()}
        </Grid>
    </>
}