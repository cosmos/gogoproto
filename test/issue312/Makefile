regenerate:
	protoc --proto_path=.:../../../../.. \
	--gogo_out=. issue312.proto 
	(cd events && make regenerate)
