import { CircularProgress, Grid, Typography } from "@mui/material";

export default function NonSuccess(props) {

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

    const renderState = () => {
        if (props.state.loading) {
            return renderNonSuccess(<CircularProgress size={25} />)
        } else if (props.state.error !== "") {
            return renderNonSuccess(
                <Typography variant="body1" color="error">
                    {props.state.error}
                </Typography>
            )
        }
        return <></>
    }

    return renderState()
}