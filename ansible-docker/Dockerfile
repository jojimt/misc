FROM centos:7
RUN yum -y install python-setuptools gcc python-devel openssl-devel
RUN easy_install pip
RUN pip install ansible==2.2.0.0
CMD /bin/bash
