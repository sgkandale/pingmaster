import { Box, Button, CircularProgress, Grid, Typography } from "@mui/material";
import { useNavigate } from 'react-router-dom';

export default function Buttons(props) {
    const navigate = useNavigate()

    return <Grid
        container
        direction="row"
        justifyContent="space-between"
        alignItems="center"
    >
        <Typography
            variant="subtitle2"
            color="error"
        >
            {props.stat.error}
        </Typography>
        <Box sx={{ mt: 8, float: 'right' }}>
            <Button
                color="secondary"
                sx={{ mr: 3 }}
                onClick={() => navigate('/targets')}
            >
                Cancel
            </Button>
            <Button
                variant="contained"
                type="submit"
                onClick={props.handleFormSubmit}
                disabled={props.stat.loading}
            >
                {props.stat.loading ? <CircularProgress size={25} /> : 'Add'}
            </Button>
        </Box>
    </Grid>
}