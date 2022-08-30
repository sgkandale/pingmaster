import { Card, CardContent, Typography } from "@mui/material";

export default function HealthStat() {
    return <Card variant="outlined">
        <CardContent>
            <Typography variant="h6">
                Health Stats
            </Typography>
            <CardContent>
                <Typography
                    variant="body1"
                >
                    <strong>Total :</strong> 10
                </Typography>
                <Typography
                    variant="body1"
                    style={{ color: 'grey' }}
                >
                    <strong>Healthy :</strong>
                </Typography>
            </CardContent>
        </CardContent>
    </Card>
}