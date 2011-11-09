package help
/* 
#include <libxml/xmlversion.h> 
#include <libxml/parser.h> 
#include <libxml/xmlstring.h> 
char* xmlChar2C(xmlChar* x) { return (char *) x; }
xmlChar* C2xmlChar(char* x) { return (xmlChar *) x; }
*/
import "C"

func XmlCheckVersion() int {
	var v C.int
	C.xmlCheckVersion(v)
	return int(v)
}

func XmlCleanUpParser() {
	C.xmlCleanupParser()
}

func XmlMemoryAllocation() int {
    return (int)(C.xmlMemBlocks())
}
