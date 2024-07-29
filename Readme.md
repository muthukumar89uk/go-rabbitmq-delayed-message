## To Run the rabbitMq server 
  # latest RabbitMQ 3.13
  - docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.13-management

## Steps to enable the delayed message scheduing plugin on rabbitMQ
- *step 1:Download the plugin*
    Download .ez file from https://github.com/rabbitmq/rabbitmq-delayed-message-exchange/releases 

- *step 2: Add the plugin to the docker image*
   RUN docker cp <path-to-the-downloaded-file> rabbitmq:/plugins/

- *step 3: Open the rabbimq bash*
    RUN docker exec -it rabbitmq /bin/bash

- *step 4: Enable the plugin*
   RUN rabbitmq-plugins enable rabbitmq_delayed_message_exchange

- *step 5: Exit from the bash*
   RUN exit

- *step 6: Restart the rabbitmq container*
   RUN docker restart rabbitmq

- *step 7: Ensure the plugin is enabled or not*
   RUN docker exec -it rabbitmq rabbitmq-plugins list

   
