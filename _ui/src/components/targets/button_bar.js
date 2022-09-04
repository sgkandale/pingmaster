import { Button, Grid } from "@mui/material"
import { AddLink, Refresh } from '@mui/icons-material'
import { useNavigate } from 'react-router-dom'

export default function ButtonBar(props) {
    const navigate = useNavigate()

    return <Grid
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
        <Button
            variant="outlined"
            startIcon={<Refresh />}
            onClick={props.fetchTargets}
            sx={{ ml: 3 }}
        >
            Refresh
        </Button>
    </Grid>
}