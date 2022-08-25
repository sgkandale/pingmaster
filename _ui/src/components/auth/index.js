import { useState } from "react"
import {
	Box, TextField, Grid, IconButton, OutlinedInput, InputLabel,
	InputAdornment, FormControl, Button, Typography, CircularProgress
} from "@mui/material";
import { Visibility, VisibilityOff } from '@mui/icons-material';

export default function Auth() {
	const [view, setView] = useState('Login')
	const [values, setValues] = useState({
		name: '',
		password: '',
		showPassword: false,
	});
	const [state, setState] = useState({
		loading: false,
		error: ""
	})


	const handleChange = (prop) => (event) => {
		setValues({ ...values, [prop]: event.target.value });
	};

	const handleClickShowPassword = () => {
		setValues({
			...values,
			showPassword: !values.showPassword,
		});
	};

	const handleMouseDownPassword = (event) => {
		event.preventDefault();
	};

	const handleFormSubmit = (event) => {
		event.preventDefault();
	}


	return <Grid
		container
		direction="column"
		justifyContent="flex-start"
		alignItems="center"
		sx={{
			height: '100vh'
		}}
	>
		<Typography
			variant="h3"
			sx={{
				fontFamily: 'Silkscreen, Helvetica, sans-serif',
				padding: '50px 0',
				color: 'primary.main',
				textTransform: 'lowercase'
			}}>
			pingmaster
		</Typography>
		<Box
			sx={{
				border: '1px solid',
				borderColor: 'text.disabled',
				borderRadius: '5px',
				minWidth: '400px',
				maxWidth: '600px',
				width: '100%',
				padding: '30px'
			}}
		>
			<Typography
				variant="h5"
				align="center"
				sx={{
					marginBottom: '10px'
				}}
			>
				{view}
			</Typography>
			<Grid
				container
				direction="column"
				justifyContent="center"
				alignItems="center"
				sx={{
					padding: '20px'
				}}
				component="form"
				onSubmit={handleFormSubmit}
			>
				<TextField
					id="outlined-basic"
					label="Name"
					variant="outlined"
					color="text"
					fullWidth
					value={values.name}
					onChange={handleChange('name')}
					sx={{
						marginBottom: '20px',
					}}
				/>
				<FormControl
					fullWidth
					variant="outlined"
					color="text"
				>
					<InputLabel
						htmlFor="outlined-adornment-password"
					>
						Password
					</InputLabel>
					<OutlinedInput
						id="outlined-adornment-password"
						type={values.showPassword ? 'text' : 'password'}
						value={values.password}
						onChange={handleChange('password')}
						fullWidth
						endAdornment={
							<InputAdornment position="end">
								<IconButton
									aria-label="toggle password visibility"
									onClick={handleClickShowPassword}
									onMouseDown={handleMouseDownPassword}
									edge="end"
								>
									{values.showPassword ? <VisibilityOff /> : <Visibility />}
								</IconButton>
							</InputAdornment>
						}
						label="Password"
					/>
				</FormControl>
				<Grid
					container
					direction="row"
					justifyContent="space-between"
					alignItems="center"
					sx={{
						marginTop: '50px'
					}}
				>
					<Button
						variant="outlined"
						onClick={() => {
							if (view === "Login") {
								setView("Register")
							} else {
								setView("Login")
							}
						}}
					>
						{view === "Login" ? "Register" : "Login"}
					</Button>
					<Button
						variant="contained"
						type="submit"
						onSubmit={handleFormSubmit}
						disabled={state.loading}
					>
						{
							state.loading ?
								<CircularProgress size={25} /> : view
						}
					</Button>
				</Grid>
			</Grid>
		</Box>
	</Grid>
}