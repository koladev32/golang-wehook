# Use an official Python runtime as a parent image
FROM python:3.11-alpine

# Set the working directory inside the container
WORKDIR /usr/src/app

# Install necessary packages and Flask
COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

# Copy the current directory contents into the container at /usr/src/app
COPY . .

# Make port 8000 available to the world outside this container
EXPOSE 8000

# Run app.py when the container launches
CMD ["python", "app.py"]
