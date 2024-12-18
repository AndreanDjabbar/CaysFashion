import axios from "axios";

export const RegisterService = async({
    dataForm,
    setIsSubmitting,
    setErrorToast,
    setSuccessToast,
    setDataForm,
    setErrors,
    navigate
}) => {
    setIsSubmitting(true);
    setErrorToast(""); 
    setSuccessToast(""); 

    try {
        const response = await axios.post("http://localhost:8080/caysfashion/register", {
            username: dataForm.username,
            email: dataForm.email,
            password: dataForm.password
        });

        setSuccessToast("Registration successful! You can now log in.");
        
        setDataForm({
            username: "",
            email: "",
            password: "",
            confirmPassword: ""
        });

        setErrors({});
        const {data} = response.data;
        setTimeout(() => navigate(`/otp-verification?randomToken=${data.randomToken}`), 2000);
    } catch (error) {
        const errorData = error.response?.data;

        if (errorData && errorData.errors) {
            const errorResponses = errorData.errors;

            let errorMessages = "Registration failed: \n";

            for (const field in errorResponses) {
                if (Object.hasOwnProperty.call(errorResponses, field)) {
                    errorMessages += `${errorResponses[field]}, `;
                }
            }
        
            setErrorToast(errorMessages);

            for (const field in errorResponses) {
                if (Object.hasOwnProperty.call(errorResponses, field)) {
                    setErrors((prevErrors) => ({
                        ...prevErrors,
                        [field]: errorResponses[field],
                    }));
                }
            }
        } else {
            setErrorToast("Unexpected error occurred. Please try again later.");
        }
    } finally {
        setIsSubmitting(false);
    }
}