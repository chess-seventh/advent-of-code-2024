"""Unit test"""
import pytest
from modules.aoc_2024.main import print_the_title
from modules.aoc_2024.exceptions.custom_error import CustomError


class TestAoc2024:
    """📦 Test class for the dummy module"""

    def test_should_test_the_dummy(self):
        """🧪 should test the dummy"""
        assert print_the_title() == "2024-11-27_aoc_2024"

    def test_should_raise_the_custom_error(self):
        """🧪 should raise the custom error"""
        with pytest.raises(CustomError):
            raise CustomError("The test Raises a Custom Exception")

