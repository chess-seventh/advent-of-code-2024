"""Custom errors module."""


class CustomError(Exception):
    """Exception raised for Custom errors."""

    def __init__(self, message="Custom error"):
        self.message = message
        super().__init__(self.message)

