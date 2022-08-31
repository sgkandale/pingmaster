import { Card, CardContent, Grid, Typography } from "@mui/material";
import { Link } from "react-router-dom";

export default function NotFound() {
    return <Grid
        container
        direction="row"
        justifyContent="center"
        alignItems="center"
        sx={{ height: '100vh' }}
    >
        <Card
            variant="outlined"
            sx={{
                padding: '50px'
            }}
        >
            <CardContent>
                <Grid
                    container
                    direction="column"
                    justifyContent="space-between"
                    alignItems="center"
                >
                    <Typography variant="h3" color="textSecondary">
                        Not Found
                    </Typography>
                    <br />
                    <Typography
                        variant="body2"
                        color="textSecondary"
                        component={Link}
                        to="/"
                    >
                        Back to Dashboard
                    </Typography>
                </Grid>
            </CardContent>
        </Card>
    </Grid>
}