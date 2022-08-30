import { Grid, Card, CardContent } from "@mui/material";
import HealthStat from "./health_stat";

export default function Overview() {
    return <Grid container spacing={2}>
        <Grid item xs={12} sm={12} md={6} lg={4} >
            <HealthStat />
        </Grid>
    </Grid>

}