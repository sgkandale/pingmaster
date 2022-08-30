import { Snackbar, Alert } from '@mui/material';

export default function LogoutSnackbar(props) {
    return <Snackbar
        open={props.open}
        autoHideDuration={3000}
        onClose={props.close}
    >
        <Alert onClose={props.close} severity={props.severity} sx={{ width: '100%' }}>
            {props.message}
        </Alert>
    </Snackbar>
}
