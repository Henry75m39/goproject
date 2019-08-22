package assembler

//This file using for implementing  application assembly.

//For example:
//Cancel a project at wrapper layer now need to call  'cancelproject' API, it requires project Id as input parameter, actually, it's really unfriendly.
//So, I do suggest in assembler layer, to implement the logic that input project name from user. Then, according to project name to
//fetch project Id through getProjectByName API. and finally.
