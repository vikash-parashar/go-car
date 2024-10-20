# go-car



gevops tutorial


firest create two different instances for aws of type aws linux

one for golang server and another one for jenkins

jenkins requires java to run

then

To add Java to the `PATH` on your AWS Linux 2 instance, follow these steps:

### **1. Verify Java Installation Path**
First, confirm the path where Java is installed:
```bash
readlink -f $(which java)
```

This command should give you something like `/usr/lib/jvm/java-11-amazon-corretto/bin/java`.

### **2. Add Java to the PATH**

1. **Open the `.bash_profile`** file in your home directory:
   ```bash
   nano ~/.bash_profile
   ```

2. **Add the following line** to set the `JAVA_HOME` environment variable and update the `PATH`:
   ```bash
   export JAVA_HOME=/usr/lib/jvm/java-11-amazon-corretto
   export PATH=$JAVA_HOME/bin:$PATH
   ```

   Make sure the path matches the output from the previous step.

3. **Save and exit** the file (if using `nano`, press `CTRL + X`, then `Y`, and `Enter`).

4. **Apply the changes to the current session**:
   ```bash
   source ~/.bash_profile
   ```

### **3. Verify Java is in the PATH**

To verify that Java has been successfully added to your `PATH`, run:

```bash
java -version
```

It should display the version of Java installed (e.g., Java 11), confirming that the `PATH` was updated correctly.

Let me know if this works!



It seems that even after importing the GPG key, you're still encountering the `GPG check FAILED` error. This is likely due to the repository or the system not fully recognizing the GPG key, or there might be some issues with the cache or system configuration.

Let's try the following steps to fix this issue:

### **1. Remove the Jenkins Repository and Re-add**

First, remove the existing Jenkins repository file:

```bash
sudo rm /etc/yum.repos.d/jenkins.repo
```

Then, re-add the repository and GPG key:

```bash
sudo wget -O /etc/yum.repos.d/jenkins.repo https://pkg.jenkins.io/redhat-stable/jenkins.repo
sudo rpm --import https://pkg.jenkins.io/redhat-stable/jenkins.io.key
```

### **2. Reinstall Jenkins**

Try reinstalling Jenkins again with the GPG check disabled:

```bash
sudo yum install --nogpgcheck jenkins -y
```

This should bypass the GPG check if it's continuing to cause issues.

### **3. Check Jenkins Status**

Once installed, start the Jenkins service:

```bash
sudo systemctl start jenkins
```

To ensure it starts on boot, enable Jenkins:

```bash
sudo systemctl enable jenkins
```

Check if the Jenkins service is running:

```bash
sudo systemctl status jenkins
```

### **4. Verify Jenkins Installation**

Once Jenkins is running, you can access it via `http://<Jenkins_EC2_Public_IP>:8080`.

Let me know if this resolves the issue, and feel free to share any output if the problem persists.





To set a password for the default `ec2-user` (or any other user) on your AWS EC2 Linux instance, follow these steps:

### **1. SSH into Your EC2 Instance**

First, connect to your EC2 instance using SSH:
```bash
ssh -i "your-key.pem" ec2-user@<EC2_Public_IP>
```

### **2. Set a Password for the User**

Use the `passwd` command to set a password for `ec2-user` (or any other user):
```bash
sudo passwd ec2-user
```

You will be prompted to enter the new password and confirm it:
```
Changing password for user ec2-user.
New password: 
Retype new password:
```



To install Golang on your AWS EC2 instance running Amazon Linux 2, follow these steps:

### **1. SSH into Your EC2 Instance**
Use the SSH command to access your EC2 instance:
```bash
ssh -i "your-key.pem" ec2-user@<Golang_EC2_Public_IP>
```

### **2. Download the Latest Golang Version**
First, visit the official [Go Downloads page](https://golang.org/dl/) and check the latest stable version. You can modify the link below with the latest version. In this example, we'll install Go 1.20.2.

Run the following command to download Golang:
```bash
wget https://golang.org/dl/go1.20.2.linux-amd64.tar.gz
```

### **3. Extract the Downloaded Archive**
Once the download is complete, extract the Golang tarball to the `/usr/local` directory (which is the recommended location for Go):

```bash
sudo tar -C /usr/local -xzf go1.20.2.linux-amd64.tar.gz
```

### **4. Set Up Go Environment Variables**
You need to add the Go binary path to your shell’s `PATH` environment variable so that you can run the `go` command from any directory.

Open or create the `.bash_profile` file:

```bash
nano ~/.bash_profile
```

Add the following lines to the file:
```bash
export PATH=$PATH:/usr/local/go/bin
```

Save and exit the file (in `nano`, press `CTRL + X`, then `Y`, and `Enter`).

### **5. Apply the Changes**
To apply the changes, source the `.bash_profile` file:

```bash
source ~/.bash_profile
```

### **6. Verify the Installation**
Check if Golang has been installed correctly by running the following command:

```bash
go version
```

Add this to your `.bash_profile`:
```bash
export GOPATH=$HOME/go
```

Source the profile again:
```bash
source ~/.bash_profile
```

---

Now you have Go installed and running on your EC2 instance! Let me know if you run into any issues.



To install Git on your Jenkins EC2 instance, follow these steps:

### Step 1: Connect to Your Jenkins EC2 Instance
If you're not already connected to your Jenkins EC2 instance, SSH into it:

```bash
ssh -i /path/to/your/key.pem ec2-user@<jenkins-ec2-public-ip>
```

(Replace `/path/to/your/key.pem` with the actual path to your PEM file and `<jenkins-ec2-public-ip>` with the public IP address of your Jenkins EC2 instance.)

### Step 2: Update the System Packages
Before installing Git, ensure that all packages on your EC2 instance are up to date:

```bash
sudo yum update -y   # for Amazon Linux/RedHat-based systems
```
### Step 3: Install Git
```bash
sudo yum install git -y
```
#### to get git version to confirm git is installed or not
```bash
git --version
```

This should display the installed version of Git.

Now, Git is installed and ready to be used on your Jenkins EC2 instance.


### install jenkins on aws linux ec2
```bash
    sudo wget -O /etc/yum.repos.d/jenkins.repo \
        https://pkg.jenkins.io/redhat/jenkins.repo
    sudo rpm --import https://pkg.jenkins.io/redhat/jenkins.io-2023.key
    sudo yum upgrade
    # Add required dependencies for the jenkins package
    sudo yum install fontconfig java-17-openjdk
    sudo yum install jenkins
```



Start Jenkins
### You can enable the Jenkins service to start at boot with the command:
```bash
    sudo systemctl enable jenkins
```
### You can start the Jenkins service with the command:
```bash
    sudo systemctl start jenkins
```
###  You can check the status of the Jenkins service using the command:
```bash
    sudo systemctl status jenkins
```




## To install java to run jenkins


Since you are using AWS EC2 with Amazon Linux, the package manager will be `yum`, and the command to install the latest version of Java will be a bit different. Here’s how you can install the latest version of Java on Amazon Linux.

### Steps to Install Latest Java on Amazon Linux (Amazon Linux 2 or Amazon Linux AMI):

### Update your system packages:**
   ```bash
   sudo yum update -y
   ```

### Install Amazon Corretto (OpenJDK Distribution by AWS):**
### Amazon Corretto is a production-ready distribution of OpenJDK. The latest version of Java provided by Amazon Corretto is Java 17. To install it, follow the steps below:

   ```bash
   sudo yum install java-17-amazon-corretto -y
   ```

### Verify the Java installation:**
###  After installation, you can verify the installed Java version with the following command:
   ```bash
   java -version
   ```

### Set Java Alternatives (if needed):**
###  If you have multiple versions of Java installed and want to switch between them, you can use the `alternatives` command to configure the default Java version:
   ```bash
   sudo alternatives --config java
   ```

Add these two setting into ec2 security groups to see jenkins on browser
```bash

Name
	
Security group rule ID
	
IP version
	
Type
	
Protocol
	
Port range
	
Source
	
Description

Name
	
Security group rule ID
	
IP version
	
Type
	
Protocol
	
Port range
	
Source
	
Description

–
sgr-076f2f56e4f4c90ce
IPv4
SSH
TCP
22
0.0.0.0/0
–
–
sgr-03efe2c41ffbe6bc8
IPv4
Custom TCP
TCP
8080
0.0.0.0/0
```

## Unlock Jenkins
### To ensure Jenkins is securely set up by the administrator, a password has been written to the log (not sure where to find it?) and this file on the server:

/var/lib/jenkins/secrets/initialAdminPassword