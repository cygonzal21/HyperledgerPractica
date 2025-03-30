A CONTINUACIÓN DESCRIBO LOS PASOS QUE REALICÉ DE LA PRÁCTICA

1. AGREGAR UN PEER A CADA ORGANIZACIÓN

   #En esta práctica modifiqué los archivos descritos abajo, con los comentarios descritos

test-network/organizations/cryptogen/crypto-config-org1.yaml
#En PeerOrgs→ Template → cambiar Count a 2 para tener dos peers por organización

test-network/organizations/cryptogen/crypto-config-org2.yaml
#En PeerOrgs→ Template → cambiar Count a 2 para tener dos peers por organización

test-network/organizations/ccp-template.yaml
#Añadir en cada item peers la definición que corresponda a un peer adicional
#línea 13
-peer1.org${ORG}.example.com

    #agregar al final línea 26. Tener cuidado con la identación:
    peer1.org${ORG}.example.com:
    url: grpcs://localhost:${P1PORT}
    tlsCACerts:
    pem: |
    ${PEERPEM}
    grpcOptions:
    ssl-target-name-override: peer1.org${ORG}.example.com
    hostnameOverride: peer1.org${ORG}.example.com

test-network/organizations/ccp-template.json
#Agregar en la línea 18
"peer1.org${ORG}.example.com"

    #Agregar en la línea 37 lo sguiente, teniendo cuidado con la identación
    "peer1.org${ORG}.example.com": {
    "url": "grpcs://localhost:${P1PORT}",
    "tlsCACerts": {
    "pem": "${PEERPEM}"
    },
    "grpcOptions": {
    "ssl-target-name-override": "peer1.org${ORG}.example.com",
    "hostnameOverride": "peer1.org${ORG}.example.com"
    }
    }

test-network/organizations/ccp-generate.sh
#Agregar un nuevo puerto para el nuevo peer de cada org
#Línea 11 y 22 agregar:
-e "s/\${P1PORT}/$6/" \
 #Agregar en la Línea 31:
P1PORT=7055
#Agregar en la Línea 38 y 39, 48 y 49 al final:
P1PORT
Agregar en la Línea 40:
P1PORT=9055

test-network/compose/compose-test-net.yaml
#Agregar en la Línea 12 los peers
peer1.org1.example.com:
peer1.org2.example.com:

    #Agregar en la Línea 99 una nueva clave para el peer1 de la org1.  FABRIC_LOGGING_SPEC #cambia a DEBUG. Tener cuidado con la identación:
    peer1.org1.example.com:
    container_name: peer1.org1.example.com
    image: hyperledger/fabric-peer:latest
    labels:
    service: hyperledger-fabric
    environment:
    - FABRIC_CFG_PATH=/etc/hyperledger/peercfg
    - FABRIC_LOGGING_SPEC=DEBUG
    #- FABRIC_LOGGING_SPEC=DEBUG
    - CORE_PEER_TLS_ENABLED=true
    - CORE_PEER_PROFILE_ENABLED=false
    - CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/fabric/tls/server.crt
    - CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/fabric/tls/server.key
    - CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/fabric/tls/ca.crt
    # Peer specific variables
    - CORE_PEER_ID=peer1.org1.example.com
    - CORE_PEER_ADDRESS=peer1.org1.example.com:7055
    - CORE_PEER_LISTENADDRESS=0.0.0.0:7055
    - CORE_PEER_CHAINCODEADDRESS=peer1.org1.example.com:7052
    - CORE_PEER_CHAINCODELISTENADDRESS=0.0.0.0:7052
    - CORE_PEER_GOSSIP_BOOTSTRAP=peer1.org1.example.com:7055
    - CORE_PEER_GOSSIP_EXTERNALENDPOINT=peer1.org1.example.com:7055
    - CORE_PEER_LOCALMSPID=Org1MSP
    - CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/fabric/msp
    - CORE_OPERATIONS_LISTENADDRESS=peer1.org1.example.com:9454
    - CORE_METRICS_PROVIDER=prometheus
    - CHAINCODE_AS_A_SERVICE_BUILDER_CONFIG={"peername":"peer1org1"}
    - CORE_CHAINCODE_EXECUTETIMEOUT=300s
    volumes:
    - ../organizations/peerOrganizations/org1.example.com/peers/peer1.org1.example.com:/etc/hyperledger/fabric
    - peer1.org1.example.com:/var/hyperledger/production
    working_dir: /root
    command: peer node start
    ports:
    - 7055:7055
    - 9454:9454
    networks:
    - test

    #Agregar en la Línea 138 una nueva clave para el peer1 de la org2. FABRIC_LOGGING_SPEC
    #se mantiene INFO. Tener cuidado con la identación:
    peer1.org2.example.com:
    container_name: peer1.org2.example.com
    image: hyperledger/fabric-peer:latest
    labels:
    service: hyperledger-fabric
    environment:
    - FABRIC_CFG_PATH=/etc/hyperledger/peercfg
    - FABRIC_LOGGING_SPEC=INFO
    #- FABRIC_LOGGING_SPEC=DEBUG
    - CORE_PEER_TLS_ENABLED=true
    - CORE_PEER_PROFILE_ENABLED=false
    - CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/fabric/tls/server.crt
    - CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/fabric/tls/server.key
    - CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/fabric/tls/ca.crt
    # Peer specific variables
    - CORE_PEER_ID=peer1.org2.example.com
    - CORE_PEER_ADDRESS=peer1.org2.example.com:9055
    - CORE_PEER_LISTENADDRESS=0.0.0.0:9055
    - CORE_PEER_CHAINCODEADDRESS=peer1.org2.example.com:9052
    - CORE_PEER_CHAINCODELISTENADDRESS=0.0.0.0:9052
    - CORE_PEER_GOSSIP_EXTERNALENDPOINT=peer1.org2.example.com:9055
    - CORE_PEER_GOSSIP_BOOTSTRAP=peer1.org2.example.com:9055
    - CORE_PEER_LOCALMSPID=Org2MSP
    - CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/fabric/msp
    - CORE_OPERATIONS_LISTENADDRESS=peer1.org2.example.com:9455
    - CORE_METRICS_PROVIDER=prometheus
    - CHAINCODE_AS_A_SERVICE_BUILDER_CONFIG={"peername":"peer1org2"}
    - CORE_CHAINCODE_EXECUTETIMEOUT=300s
    volumes:
    - ../organizations/peerOrganizations/org2.example.com/peers/peer1.org2.example.com:/etc/hyperledger/fabric
    - peer1.org2.example.com:/var/hyperledger/production
    working_dir: /root
    command: peer node start
    ports:
    - 9055:9055
    - 9455:9455
    networks:
    - test

test-network/compose/compose-couch.yaml
#Agregar en la Línea 41 una nueva clave para el peer1 de la org1, manteniendo el puerto 5984. Tener cuidado con la identación:
Couchdb01:
container_name: couchdb01
image: couchdb:3.4.2
labels:
service: hyperledger-fabric
environment: - COUCHDB_USER=admin - COUCHDB_PASSWORD=adminpw
networks: - test

    peer1.org1.example.com:
    environment:
    - CORE_LEDGER_STATE_STATEDATABASE=CouchDB
    - CORE_LEDGER_STATE_COUCHDBCONFIG_COUCHDBADDRESS=couchdb01:5984
    - CORE_LEDGER_STATE_COUCHDBCONFIG_USERNAME=admin
    - CORE_LEDGER_STATE_COUCHDBCONFIG_PASSWORD=adminpw
    depends_on:
    - couchdb01

    #Agregar en la Línea 70 una nueva clave para el peer1 de la org2,  manteniendo el puerto 5984.Tener cuidado con la identación:
    couchdb11:
    container_name: couchdb11
    image: couchdb:3.4.2
    labels:
    service: hyperledger-fabric
    environment:
    - COUCHDB_USER=admin
    - COUCHDB_PASSWORD=adminpw
    networks:
    - test

    peer1.org2.example.com:
    environment:
    - CORE_LEDGER_STATE_STATEDATABASE=CouchDB
    - CORE_LEDGER_STATE_COUCHDBCONFIG_COUCHDBADDRESS=couchdb11:5984
    - CORE_LEDGER_STATE_COUCHDBCONFIG_USERNAME=admin
    - CORE_LEDGER_STATE_COUCHDBCONFIG_PASSWORD=adminpw
    depends_on:
    - couchdb11

test-network/compose/docker/docker-compose-test-net.yaml
#Agregar en la Línea 21 una nuevo servicio para el peer1 de la org1. Tener cuidado con la identación:
peer1.org1.example.com:
container_name: peer1.org1.example.com
image: hyperledger/fabric-peer:latest
labels:
service: hyperledger-fabric
environment:
#Generic peer variables - CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock - CORE_VM_DOCKER_HOSTCONFIG_NETWORKMODE=fabric_test
volumes: - ./docker/peercfg:/etc/hyperledger/peercfg - ${DOCKER_SOCK}:/host/var/run/docker.sock
#Agregar en la Línea 33 una nuevo servicio para el peer1 de la org2.Tener cuidado con la identación:
peer1.org2.example.com:
container_name: peer1.org2.example.com
image: hyperledger/fabric-peer:latest
labels:
service: hyperledger-fabric
environment:
#Generic peer variables - CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock - CORE_VM_DOCKER_HOSTCONFIG_NETWORKMODE=fabric_test
volumes: - ./docker/peercfg:/etc/hyperledger/peercfg - ${DOCKER_SOCK}:/host/var/run/docker.sock

2. REALICÉ UN CHAINCODE LLAMADO product.go, el cual tiene la siguiente lógica de negocio.

   a. Crea los assets claves para registrar un producto en la red
   b. El fabricante (Org1) registra/crea un producto en la red con estado "fabricado".
   c. El transportista (Org2) recibe el producto y cambia su estado a "en transito".
   d. El distribuidor (Org3) recibe el producto y cambia su estado a "Entregado al minorista".
   NOTA: El archivo está en la ruta ../fabric-samples/asset-transfer-basic/chaincode-go.
   Tuve inconvenientes para desplegarlo :(

3. TUMBAR Y LEVANTAR LA RED
   ./network.sh down
   ./network.sh up createChannel -s couchdb

4. UNIR LOS NUEVOS PEERS AL CANAL
   #peer1 de la org1
   export PATH=${PWD}/../bin:$PATH
   export FABRIC_CFG_PATH=${PWD}/../config/ 
    export CORE_PEER_TLS_ENABLED=true 
    export CORE_PEER_LOCALMSPID="Org1MSP" 
    export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/ca.crt
   export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp export CORE_PEER_ADDRESS=localhost:7055
   peer channel join -b channel-artifacts/mychannel.block

   #peer1 de la org2
   export PATH=${PWD}/../bin:$PATH
   export FABRIC_CFG_PATH=${PWD}/../config/ 
    export CORE_PEER_TLS_ENABLED=true 
    export CORE_PEER_LOCALMSPID="Org2MSP" 
    export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls/ca.crt
   export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp export CORE_PEER_ADDRESS=localhost:9055
   peer channel join -b channel-artifacts/mychannel.block

5. AÑADIR UNA NUEVA ORGANIZACIÓN
   cd addOrg3
   ./addOgr3.sh up -s couchdb

6. EMPAQUETAR EL CHAINCODE (No me funcionó pero esta sería la instrucción)
   peer lifecycle chaincode package product.tar.gz --path . --lang golang --label product_1.0

7. INSTALAR EL CHAINCODE EN TODOS LOS PEERS
   #Para Org1 (peer0 y peer1)
   export CORE_PEER_LOCALMSPID="Org1MSP"
   export CORE_PEER_MSPCONFIGPATH=${PWD}/../organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
    export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/../organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
   export CORE_PEER_ADDRESS=localhost:7051

   peer lifecycle chaincode install product.tar.gz

   export CORE_PEER_ADDRESS=localhost:7055
   peer lifecycle chaincode install product.tar.gz

   #Para Org2 (peer0 y peer1)
   export CORE_PEER_LOCALMSPID="Org2MSP"
   export CORE_PEER_MSPCONFIGPATH=${PWD}/../organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
    export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/../organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
   export CORE_PEER_ADDRESS=localhost:9051

   peer lifecycle chaincode install product.tar.gz

   export CORE_PEER_ADDRESS=localhost:9055
   peer lifecycle chaincode install product.tar.gz

   #Para Org3 (único peer)
   export CORE_PEER_LOCALMSPID="Org3MSP"
   export CORE_PEER_MSPCONFIGPATH=${PWD}/../organizations/peerOrganizations/org3.example.com/users/Admin@org3.example.com/msp
    export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/../organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt
   export CORE_PEER_ADDRESS=localhost:11051

   peer lifecycle chaincode install product.tar.gz

8. COMMITEAR EL CHAINCODE
   peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID mychannel \
   --name product --version 1.0 --sequence 1 --tls true --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
   --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt \
   --peerAddresses localhost:7055 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/ca.crt \
   --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt \
   --peerAddresses localhost:9055 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls/ca.crt \
   --peerAddresses localhost:11051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt

9. INVOCAR EL CHAINCODE
   #Crear un producto (solo Org1 puede)
   peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls true \
   --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
   -C mychannel -n product --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt \
   --peerAddresses localhost:7055 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/ca.crt \
   --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt \
   --peerAddresses localhost:9055 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls/ca.crt \
   --peerAddresses localhost:11051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt \
   --isInit -c '{"Args":["CreateProduct","P123","Laptop"]}'
