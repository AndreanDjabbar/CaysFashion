import { useEffect, useState } from "react";
import { AuthLayout } from "../layouts/AuthLayout";
import { AuthInput } from "../fragments/AuthInput";
import { Button } from "../elements/Button";
import { AuthFormLayout } from "../layouts/AuthFormLayout";
import { FooterLayout } from "../layouts/FooterLayout";
import { Link } from "../elements/Link";
import { NavLayout } from "../layouts/NavLayout";
import { ErrorToaster, SuccessToaster } from "../elements/Toast";
import "../../styles/Register.css";
import loading from "../../assets/loading.png";
import { useNavigate } from "react-router-dom";
import { RegisterService as registerHandler } from "../../services/RegisterService";

export const RegisterPage = () => {
    const navigate = useNavigate();
    const [dataForm, setDataForm] = useState({
        username: "",
        email: "",
        password: "",
        confirmPassword: ""
    });
    const [isSubmitting, setIsSubmitting] = useState(false);
    const [errors, setErrors] = useState({});
    const [isFilled, setIsFilled] = useState(false);
    const [errorToast, setErrorToast] = useState(""); 
    const [successToast, setSuccessToast] = useState("");

    const validateField = (name, value) => {
        switch (name) {
            case "username":
                return value.length < 6 ? "Username needs to be 6 characters or more" : "";
            case "email":
                return !/\S+@\S+\.\S+/.test(value) ? "Email is invalid" : "";
            case "password":
                return value.length < 8 ? "Password needs to be 8 characters or more" : "";
            case "confirmPassword":
                return value !== dataForm.password ? "Passwords do not match" : "";
            default:
                return "";
        }
    };

    const handleChange = (e) => {
        const { name, value } = e.target;
        const updatedErrors = { ...errors, [name]: validateField(name, value) };
        
        setErrors(updatedErrors);
        setDataForm({ ...dataForm, [name]: value });
    };

    useEffect(() => {
        const filled = Object.values(dataForm).every((val) => val !== "");
        const noErrors = !Object.values(errors).some((error) => error);
        setIsFilled(filled && noErrors);
    }, [dataForm, errors]);

    const onSubmit = (e) => {
        e.preventDefault();
        if (!Object.values(errors).some((error) => error)) {
            registerHandler({
                dataForm,
                setIsSubmitting,
                setErrorToast,
                setSuccessToast,
                setDataForm,
                setErrors,
                navigate
            });
        } else {
            setErrorToast("Please fix the form errors before submitting.");
        }
    };

    return (
        <AuthLayout>
            <NavLayout className="register-nav">
                <Link href="/login" id="register">Login</Link>
                <Link href="" id="explore">Explore About</Link>
            </NavLayout>

            {errorToast && <ErrorToaster>{errorToast}</ErrorToaster>}
            {successToast && <SuccessToaster>{successToast}</SuccessToaster>} 

            <AuthFormLayout
                className="register-form"
                titleForm="Register"
                guide="Already Have Account?"
                guideType="Login"
                guideLink="/login"
                onSubmit={onSubmit}
            >
                <div className="register-form-input">
                    <AuthInput
                        className="register-username"
                        id="username"
                        name="username"
                        placeholder="Username"
                        type="text"
                        value={dataForm.username}
                        onChange={handleChange}
                        errorMessage={errors.username}
                    >Username</AuthInput>

                    <AuthInput
                        className="register-email"
                        id="email"
                        name="email"
                        placeholder="Email"
                        type="email"
                        value={dataForm.email}
                        onChange={handleChange}
                        errorMessage={errors.email}
                    >Email</AuthInput>

                    <AuthInput
                        className="register-password"
                        id="password"
                        name="password"
                        placeholder="Password"
                        type="password"
                        value={dataForm.password}
                        onChange={handleChange}
                        errorMessage={errors.password}
                    >Password</AuthInput>

                    <AuthInput
                        className="register-password2"
                        id="confirmPassword"
                        name="confirmPassword"
                        placeholder="Confirm Password"
                        type="password"
                        value={dataForm.confirmPassword}
                        onChange={handleChange}
                        errorMessage={errors.confirmPassword}
                    >Confirm Password</AuthInput>
                </div>
                <div className="register-form-button">
                    <Button
                        type="submit"
                        disabled={Object.values(errors).some((error) => error) || !isFilled || isSubmitting}
                        className="register-button"
                    >
                        {isSubmitting ? (
                            <div className="loading-section">
                                <p>Registering..</p>
                                <img src={loading} alt="Loading..." className="loading-icon" />
                            </div>
                        ) : (
                            "Register"
                        )}
                    </Button>
                </div>
            </AuthFormLayout>
            <FooterLayout className="register-footer"></FooterLayout>
        </AuthLayout>
    );
};