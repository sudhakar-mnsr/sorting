/*hello_server.c*/
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <netdb.h>

#define BUF_SIZE 4096

int main(){

  char hostname[]="127.0.0.1"; 
  short port=1845; 
  struct sockaddr_in saddr_in;         // Server address
  struct sockaddr_in client_saddr_in;  // Client address
  socklen_t saddr_len;                 // Length of address
  int server_sock, peer_sock_fd;        // Socket file descriptor
  char response[BUF_SIZE];          
  int n;                             

  //set up the address information
  saddr_in.sin_family = AF_INET;
  inet_aton(hostname, &saddr_in.sin_addr);
  saddr_in.sin_port = htons(port);
  saddr_len = sizeof(struct sockaddr_in);  

  // Open a socket
  if( (server_sock = socket(AF_INET, SOCK_STREAM, 0))  < 0){
    perror("socket");
    exit(1);
  }

  // Bind the socket to its address
  if(bind(server_sock, (struct sockaddr *) &saddr_in, saddr_len) < 0){
    perror("bind");
    exit(1);
  }

  // Listen. Queue up to 5 pending connections.
  if(listen(server_sock, 5)  < 0){
    perror("listen");
    exit(1);
  }
  printf("Listening On: %s:%d\n", inet_ntoa(saddr_in.sin_addr), ntohs(saddr_in.sin_port));

  // Accept incoming connections
  if((peer_sock_fd = accept(server_sock, (struct sockaddr *) &client_saddr_in, &saddr_len)) < 0){
    perror("accept");
    exit(1);
  }

  printf("Connection From: %s:%d (%d)\n", 
         inet_ntoa(client_saddr_in.sin_addr), //address as dotted quad
         ntohs(client_saddr_in.sin_port),     //the port in host order
         peer_sock_fd);                        //the file descriptor number

  printf("Connection To: %s:%d\n", 
         inet_ntoa(saddr_in.sin_addr),        //address as dotted quad
         ntohs(saddr_in.sin_port));           //the port in host order

  // Read data from client
  if((n = read(peer_sock_fd,response, BUF_SIZE-1)) < 0){
    perror("read");
    exit(1);
  }
  response[n] = '\0';
  printf("Read from client: %s", response);

  // Construct and send response
  snprintf(response, BUF_SIZE, "Hello %s:%d \nGo Navy! Beat Army\n", 
           inet_ntoa(client_saddr_in.sin_addr),    //address as dotted quad
           ntohs(client_saddr_in.sin_port));       //the port in host order
  printf("Sending: %s",response);
  if(write(peer_sock_fd, response, strlen(response)) < 0){
    perror("write");
    exit(1);
  }

  // Close client and server sockets
  printf("Closing socket\n\n");
  close(peer_sock_fd);
  close(server_sock);

  return 0;
}
