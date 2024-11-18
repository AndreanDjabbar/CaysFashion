import { useLocation } from "react-router-dom";
import { OtpVerifLayout } from "../layouts/OtpVerifLayout";
import { OtpForm } from "../fragments/OtpForm";
import { useState, useEffect } from "react";
import "../../styles/OtpVerif.css";
import { OtpService } from "../../services/OtpService";
import { SuccessToaster, ErrorToaster } from "../elements/Toast";
import loading from "../../assets/loading.png";
import { useNavigate } from "react-router-dom";

export const OtpVerifPage = () => {
    const navigate = useNavigate();
    const location = useLocation();
    const [otp, setOtp] = useState(new Array(4).fill("")); 
    const [isSubmitting, setIsSubmitting] = useState(false);
    const [isOtpSubmitted, setIsOtpSubmitted] = useState(false); 
    const [errorToast, setErrorToast] = useState(""); 
    const [successToast, setSuccessToast] = useState(""); 
    const [error, setError] = useState("");
    const [isTokenValid, setIsTokenValid] = useState(true); 
    const randomToken = new URLSearchParams(location.search).get('randomToken');

    useEffect(() => {
        if (!randomToken) {
            navigate("/register");
        } else {
            const verifyToken = async () => {
                try {
                    const response = await fetch(`http://localhost:8080/caysfashion/verify-token?randomToken=${randomToken}`);
                    console.log(response);
                    if (response.ok == true) {
                        const data = await response.json();
                        setIsTokenValid(true);
                    } else {
                        throw new Error("Invalid token or session expired");
                    }
                } catch (error) {
                    setIsTokenValid(false); 
                    setErrorToast(error.message);
                    setTimeout(() => {
                        navigate("/register");
                    }, 2000); 
                }
            };

            verifyToken();
        }
    }, [randomToken, navigate]);

    const handleChange = (element, index) => {
        const value = element.value;
        setError("");
        if (/^[0-9]$/.test(value) || value === "") {
            const newOtp = [...otp];
            newOtp[index] = value;
            setOtp(newOtp);

            if (value && element.nextSibling) {
                element.nextSibling.focus();
            }
        }
    };

    const handleBackspace = (element, index) => {
        const newOtp = [...otp];
        newOtp[index] = ""; 
        setOtp(newOtp);

        if (index > 0 && !element.value && element.previousSibling) {
            element.previousSibling.focus();
        }
    };

    useEffect(() => {
        const otpCode = otp.join("");
        if (otpCode.length === 4 && !otp.includes("") && !isOtpSubmitted) {
            setIsSubmitting(true); 
            setIsOtpSubmitted(true);
            setTimeout(() => handleOtpSubmit(otpCode), 2000);
        }
    }, [otp]);

    const handleOtpSubmit = async (otpCode) => {
        try {
            const response = await OtpService(otpCode);
            setSuccessToast(response.message);
            setErrorToast(""); 
        } catch (error) {
            setErrorToast(`OTP verification failed: ${error.message}`);
            setError(error.message);
            setSuccessToast("");
            setOtp(new Array(4).fill("")); 
        } finally {
            setIsSubmitting(false);
            setIsOtpSubmitted(false);
        }
    };

    useEffect(() => {
        if (successToast) {
            setOtp(new Array(4).fill("")); 
            setError("");
            setTimeout(() => {
                navigate("/login");
            }, 3000); 
        }
    }, [successToast, navigate]);

    return (
        <OtpVerifLayout className="otp-verif-container">
            {errorToast && <ErrorToaster>{errorToast}</ErrorToaster>}
            {successToast && <SuccessToaster>{successToast}</SuccessToaster>}

            <OtpForm
                className="otp-form-container"
                title="OTP Verification"
                description="Please enter the OTP code sent to your email"
            >
                <div className="otp-input-container">
                    {otp.map((data, index) => (
                        <input
                            key={index}
                            type="text"
                            maxLength="1"
                            value={data}
                            onChange={(e) => handleChange(e.target, index)}
                            onKeyDown={(e) => {
                                if (e.key === "Backspace") {
                                    handleBackspace(e.target, index);
                                }
                            }}
                            onFocus={(e) => e.target.select()}
                            className="otp-input"
                            disabled={isSubmitting}
                        />
                    ))}
                </div>
                {error && <p id="error">{error}</p>}
                {isSubmitting && (
                    <div className="loading-section">
                        <p>Verifying OTP...</p>
                        <img src={loading} alt="Loading..." className="loading-icon" />
                    </div>
                )}
            </OtpForm>
        </OtpVerifLayout>
    );
};