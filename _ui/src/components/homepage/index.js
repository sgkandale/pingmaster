import { Grid, Card, CardContent } from "@mui/material";

export default function Homepage() {
    return <Grid container spacing={2}>
        <Grid item xs={12} sm={12} md={6} lg={4} >
            <Card variant="outlined">
                <CardContent>
                    Hello
                </CardContent>
            </Card>
        </Grid>
    </Grid>

}