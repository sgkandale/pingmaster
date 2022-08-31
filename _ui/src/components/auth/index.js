import { useState } from "react"
import {
	Box, TextField, Grid, IconButton, OutlinedInput, InputLabel,
	InputAdornment, FormControl, Button, Typography, CircularProgress
} from "@mui/material";
import { Visibility, VisibilityOff } from '@mui/icons-material';
import axios from 'axios'
import { ServerAddr } from '../server'
import { useDispatch } from 'react-redux'
import { ACTION_LOGIN } from "../state_actions";
import { useNavigate } from 'react-router-dom';

export default function Auth() {
	const [view, setView] = useState('Login')
	const [values, setValues] = useState({
		name: '',
		password: '',
		showPassword: false,
		loading: false,
		error: ""
	})
	const dispatch = useDispatch()
	const navigate = useNavigate();

	const handleChange = (prop) => (event) => {
		setValues({
			...values,
			[prop]: event.target.value,
			loading: false
		});
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
		setValues({
			...values,
			loading: true,
		})
		var addr = '/user/login'
		if (view === "Register") {
			addr = '/user/'
		}
		axios.post(ServerAddr + addr, {
			name: values.name,
			password: values.password,
		}, {
			headers: {
				'Content-Type': 'application/json',
				'Authorization': "not-auth"
			}
		})
			.then(response => {
				setValues({
					...values,
					loading: false,
				})
				dispatch({ type: ACTION_LOGIN, payload: response.data.response })
				navigate('/')
			})
			.catch(error => {
				setValues({
					...values,
					loading: false,
					error: error.response.data.message
				})
			});
	}


	return <Grid
		container
		direction="column"
		justifyContent="flex-start"
		alignItems="center"
		sx={{ height: '100vh' }}
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
				noValidate
				autoComplete="off"
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
					disabled={values.showPassword}
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
						disabled={values.showPassword}
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
						disabled={values.loading}
					>
						{
							values.loading ?
								<CircularProgress size={25} /> : view
						}
					</Button>
				</Grid>
				<Typography
					variant="body1"
					color="error"
					sx={{
						marginTop: '20px'
					}}
				>
					{values.error}
				</Typography>
			</Grid>
		</Box>
	</Grid>
}